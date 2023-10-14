package scyna_test

import (
	"log"
	"testing"
	"time"

	scyna "github.com/scyna/core"
	scyna_proto "github.com/scyna/core/proto"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type endpointTest struct {
	status             int32
	url                string
	channel            string
	request            proto.Message
	response           proto.Message
	event              proto.Message
	exactEventMatch    bool
	exactResponseMatch bool
	errorCodeOnly      bool
}

type endpointTestResult struct {
	t         *testing.T
	eventData []byte
	response  *scyna_proto.Response
}

func (result *endpointTestResult) DecodeResponse(response proto.Message) *endpointTestResult {
	if result.response == nil {
		result.t.Fatal("No response")
	}

	if proto.Unmarshal(result.response.Body, response) != nil {
		result.t.Fatal("Error in decode response")
	}

	return result
}

func (result *endpointTestResult) DecodeEvent(event proto.Message) *endpointTestResult {
	if result.eventData == nil {
		result.t.Fatal("No event")
	}

	if proto.Unmarshal(result.eventData, event) != nil {
		result.t.Fatal("Error in decode event")
	}

	return result
}

func Endpoint(url string) *endpointTest {
	return &endpointTest{
		url:           url,
		errorCodeOnly: false,
	}
}

func (t *endpointTest) WithRequest(request proto.Message) *endpointTest {
	t.request = request
	return t
}

func (t *endpointTest) ExpectError(err scyna.Error) *endpointTest {
	e_ := scyna_proto.Error{
		Code:    err.Code(),
		Message: err.Message(),
	}
	t.errorCodeOnly = false
	t.status = 400
	t.response = &e_
	return t
}

func (t *endpointTest) ExpectErrorCode(code string) *endpointTest {
	e_ := scyna_proto.Error{
		Code: code,
	}
	t.errorCodeOnly = true
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

func (t *endpointTest) ExpectResponseLike(response proto.Message) *endpointTest {
	t.status = 200
	t.response = response
	t.exactResponseMatch = false
	return t
}

func (t *endpointTest) ExpectEvent(event proto.Message) *endpointTest {
	t.event = event
	t.exactEventMatch = true
	return t
}

func (t *endpointTest) ExpectEventLike(event proto.Message) *endpointTest {
	t.event = event
	t.exactEventMatch = false
	return t
}

func (t *endpointTest) FromChannel(channel string) *endpointTest {
	t.channel = channel
	return t
}

func (st *endpointTest) Run(t *testing.T) *endpointTestResult {
	streamName := scyna.Module()
	if len(st.channel) > 0 {
		createStream(streamName)
	}

	var res = st.callEndpoint(t)
	ret := &endpointTestResult{response: res}
	if st.status != res.Code {
		t.Fatalf("Expect status %d but actually %d with response %s", st.status, res.Code, string(res.Body))
	}

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

	if st.event != nil {
		if len(st.channel) > 0 {
			subs, err := scyna.JetStream.SubscribeSync(streamName + "." + st.channel)
			if err != nil {
				t.Fatal("Error in subscribe")
			}

			msg, err := subs.NextMsg(time.Second)
			if err != nil {
				t.Fatal("Timeout")
			}

			var event scyna_proto.Event
			ret.eventData = msg.Data
			if err := proto.Unmarshal(msg.Data, &event); err != nil {
				log.Print("Register unmarshal error response data:", err.Error())
				t.Fatal("Can not parse received event")
			}

			receivedEvent := proto.Clone(st.event)
			if proto.Unmarshal(event.Body, receivedEvent) != nil {
				t.Fatal("Can not parse received event")
			}

			if st.exactEventMatch {
				if !proto.Equal(st.event, receivedEvent) {
					t.Fatal("Event not match")
				}
			} else {
				if !matchMessage(st.event, receivedEvent) {
					t.Fatal("Event not match")
				}
			}
			subs.Unsubscribe()
		} else {
			time.Sleep(time.Millisecond * 100)
			receivedEvent := nextEvent()
			if receivedEvent == nil {
				t.Fatal("No event received")
			}
			ret.eventData, _ = proto.Marshal(receivedEvent)
			if st.exactEventMatch {
				if !proto.Equal(st.event, receivedEvent) {
					t.Fatal("Event not match")
				}
			} else {
				if !matchMessage(st.event, receivedEvent) {
					t.Fatal("Event not match")
				}
			}
		}
	}

	if len(st.channel) > 0 {
		deleteStream(streamName)
	}

	return ret
}

func (st *endpointTest) callEndpoint(t *testing.T) *scyna_proto.Response {
	req := scyna_proto.Request{TraceID: scyna.ID.Next(), JSON: false}
	res := scyna_proto.Response{}

	if st.request != nil {
		var err error
		if req.Body, err = proto.Marshal(st.request); err != nil {
			t.Fatal("Bad Request:", err)
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := scyna.Nats.Request(scyna_utils.PublishURL(st.url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				t.Fatal("Server Error:", err)
			}
		} else {
			t.Fatal("Server Error:", err)
		}
	} else {
		t.Fatal("Bad Request:", err)
	}

	return &res
}
