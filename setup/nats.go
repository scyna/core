package scyna_setup

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func CreateStreamIfMissing(module string) {
	if _, err := JetStream.StreamInfo(module); err == nil {
		log.Print("Stream:", module, " exists")
		return
	}

	if _, err := JetStream.AddStream(&nats.StreamConfig{
		Name:     module,
		Subjects: []string{module + ".>"},
		Storage:  nats.FileStorage,
		MaxAge:   time.Hour * 24 * 7, //keep for a week
	}); err != nil {
		panic("Error in creating stream")
	}

	log.Print("Create stream:", module)
}

func CreateConsumerIfMissing(sender, receiver string) {
	if _, err := JetStream.StreamInfo(sender); err != nil {
		panic("No stream `" + sender + "`")
	}

	if _, err := JetStream.ConsumerInfo(sender, receiver); err == nil {
		log.Print("Consumer:", sender, " exists")
		return
	}

	if _, err := JetStream.AddConsumer(sender, &nats.ConsumerConfig{
		Durable:   receiver,
		AckPolicy: nats.AckExplicitPolicy,
	}); err != nil {
		panic("Error in creating consumer:" + err.Error())
	}
	log.Print("Create consumer:", sender, " for stream:", receiver)
}
