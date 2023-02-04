package scyna_test

import (
	"log"

	"github.com/nats-io/nats.go"
	scyna "github.com/scyna/core"
)

func deleteStream(name string) error {
	err := scyna.JetStream.DeleteStream(name)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func createStream(name string) error {
	if _, err := scyna.JetStream.StreamInfo(name); err == nil {
		scyna.JetStream.DeleteStream(name)
	}

	if _, err := scyna.JetStream.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".>"},
	}); err != nil {
		log.Print("Create stream for module " + name + ": Error: " + err.Error())
		return err
	}
	return nil
}
