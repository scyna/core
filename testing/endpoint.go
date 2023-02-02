package scyna_test

import (
	"log"
	"testing"
	"time"

	scyna "github.com/scyna/core"
	scyna_proto "github.com/scyna/core/proto/generated"
	"google.golang.org/protobuf/proto"
)

type endpointTest struct {
	status             int32
	url                string
	request            proto.Message
	response           proto.Message
	exactResponseMatch bool
	expectedEvent      proto.Message
	channel            string
	exactEventMatch    bool
}

func EndpointTest(url string) *endpointTest {
	return &endpointTest{url: url}
}

func (t *endpointTest) WithRequest(request proto.Message) *endpointTest {
	t.request = request
	return t
}

func (t *endpointTest) PublishEventTo(channel string) *endpointTest {
	t.channel = channel
	return t
}

func (t *endpointTest) ExpectError(err scyna.Error) *endpointTest {
	e_ := scyna_proto.Error{
		Code:    err.Code(),
		Message: err.Message(),
	}
	t.status = 400
	t.response = &e_
	return t
}

func (t *endpointTest) ExpectSuccess() *endpointTest {
	t.status = 200
	return t
}

func (t *endpointTest) ExpectResponse(response proto.Message) *endpointTest {
	t.status = 200
	t.response = response
	t.exactResponseMatch = true
	return t
}

func (t *endpointTest) MatchResponse(response proto.Message) *endpointTest {
	t.status = 200
	t.response = response
	t.exactResponseMatch = false
	return t
}

func (t *endpointTest) ExpectEvent(event proto.Message) *endpointTest {
	t.expectedEvent = event
	t.exactEventMatch = true
	return t
}

func (t *endpointTest) MatchEvent(event proto.Message) *endpointTest {
	t.expectedEvent = event
	t.exactEventMatch = false
	return t
}

func (st *endpointTest) Run(t *testing.T, response ...proto.Message) {

	streamName := getStreamName(st.channel)
	if len(st.channel) > 0 && len(streamName) == 0 {
		t.Fatal("Invalid channel format")
	}

	if len(streamName) > 0 {
		createStreamForModule(streamName)
	}

	var res = st.callEndpoint(t)
	if st.status != res.Code {
		t.Fatalf("Expect status %d but actually %d with response %s", st.status, res.Code, string(res.Body))
	}

	if len(response) == 0 {
		if st.response != nil {
			tmp := proto.Clone(st.response)
			if err := proto.Unmarshal(res.Body, tmp); err != nil {
				t.Fatal("Can not parse response body")
			}

			if st.exactResponseMatch {
				if !proto.Equal(tmp, st.response) {
					t.Fatal("Response not match")
				}
			} else {
				if !matchMessage(tmp, st.response) {
					t.Fatal("Response not match")
				}
			}
		}
	} else if len(response) == 1 {
		if err := proto.Unmarshal(res.Body, response[0]); err != nil {
			t.Fatal("Can not parse response body")
		}
	} else {
		t.Fatal("Too many parametter")
	}

	if st.expectedEvent != nil {
		subs, err := scyna.JetStream.SubscribeSync(streamName + ".*")
		if err != nil {
			t.Fatal("Error in subscribe")
		}

		msg, err := subs.NextMsg(time.Second)
		if err != nil {
			t.Fatal("Timeout")
		}

		var event scyna_proto.Event
		if err := proto.Unmarshal(msg.Data, &event); err != nil {
			log.Print("Register unmarshal error response data:", err.Error())
			t.Fatal("Can not parse received event")
		}

		receivedEvent := proto.Clone(st.expectedEvent)
		if proto.Unmarshal(event.Body, receivedEvent) != nil {
			t.Fatal("Can not parse received event")
		}

		if st.exactEventMatch {
			if !proto.Equal(st.expectedEvent, receivedEvent) {
				t.Fatal("Event not match")
			}
		} else {
			if !matchMessage(st.expectedEvent, receivedEvent) {
				t.Fatal("Event not match")
			}
		}

		subs.Unsubscribe()
		subs.Drain()
	}

	if len(streamName) > 0 {
		deleteStream(streamName)
	}
}

func (st *endpointTest) callEndpoint(t *testing.T) *scyna_proto.Response {
	context := scyna.Trace{
		ID:       scyna.ID.Next(),
		ParentID: 0,
		Time:     time.Now(),
		Path:     st.url,
		Type:     scyna.TRACE_ENDPOINT,
		Source:   "scyna.test",
	}
	defer context.Record()

	req := scyna_proto.Request{TraceID: context.ID, JSON: false}
	res := scyna_proto.Response{}

	if st.request != nil {
		var err error
		if req.Body, err = proto.Marshal(st.request); err != nil {
			t.Fatal("Bad Request:", err)
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := scyna.Connection.Request(scyna.PublishURL(st.url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				t.Fatal("Server Error:", err)
			}
		} else {
			t.Fatal("Server Error:", err)
		}
	} else {
		t.Fatal("Bad Request:", err)
	}

	context.SessionID = res.SessionID
	context.Status = res.Code
	return &res
}
