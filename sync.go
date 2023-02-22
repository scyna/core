package scyna

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	scyna_proto "github.com/scyna/core/proto/generated"
	scyna_utils "github.com/scyna/core/utils"
	"google.golang.org/protobuf/proto"
)

type SyncHandler[R proto.Message] func(ctx *Endpoint, data R) *http.Request

func RegisterSync[R proto.Message](channel string, receiver string, handler SyncHandler[R]) {
	subject := module + "." + channel
	durable := receiver
	Session.Info(fmt.Sprintf("Channel %s, durable: %s", subject, durable))
	event := scyna_utils.NewMessageForType[R]()

	sub, err := JetStream.PullSubscribe(subject, durable, nats.BindStream(module))

	if err != nil {
		panic("Error in start event stream:" + err.Error())
	}

	go func() {
		for {
			messages, err := sub.Fetch(1)
			if err != nil || len(messages) != 1 {

				continue
			}
			m := messages[0]

			var msg scyna_proto.Event
			if err := proto.Unmarshal(m.Data, &msg); err != nil {
				log.Print("Register unmarshal error response data:", err.Error())
				m.Ack()
				continue
			}
			trace := Trace{
				Path:      subject, //FIXME
				SessionID: Session.ID(),
				Type:      TRACE_SYNC,
				Time:      time.Now(),
				ID:        ID.Next(),
				ParentID:  msg.TraceID,
			}

			context := NewContext(trace.ID)

			if err := proto.Unmarshal(msg.Body, event); err != nil {
				log.Print("Error in parsing data:", err)
				m.Ack()
				continue
			}

			request := handler(context, event)
			if sendSyncRequest(request) {
				m.Ack()
			} else {
				sent := false
				for i := 0; i < 3; i++ {
					request := handler(context, event)
					if sendSyncRequest(request) {
						m.Ack()
						sent = true
						break
					}
					time.Sleep(time.Second * 30)
				}

				if !sent {
					m.Nak()
				}
			}
			trace.Record()
		}
	}()
}

func sendSyncRequest(request *http.Request) bool {
	if request == nil {
		return true
	}

	response, err := HttpClient().Do(request)
	if err != nil {
		Session.Warning("Sync:" + err.Error())
		return false
	} else {
		defer response.Body.Close()
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			Session.Info("Sync error: " + err.Error())
			return true
		}
		bodyString := string(bodyBytes)
		Session.Info(fmt.Sprintf("Sync: %s - %d - %s", request.URL, response.StatusCode, bodyString))

		if response.StatusCode >= 500 && response.StatusCode <= 599 {
			return false
		}
	}
	return true
}
