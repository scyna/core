package scyna_test

import (
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	scyna "github.com/scyna/core"
)

func getStreamName(channel string) string {
	list := strings.Split(channel, ".")

	if len(list) > 1 {
		return list[0]
	}

	return ""
}

func deleteStream(name string) error {
	err := scyna.JetStream.DeleteStream(name)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func createStreamForModule(module string) error {
	if _, err := scyna.JetStream.AddStream(&nats.StreamConfig{
		Name:     module,
		Subjects: []string{module + ".>"},
	}); err != nil {
		log.Print("Create stream for module " + module + ": Error: " + err.Error())
		return err
	}
	return nil
}
