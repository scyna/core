package scyna_setup

import (
	"time"

	"github.com/nats-io/nats.go"
)

func createStreamIfMissing(module string) {
	if _, err := JetStream.AddStream(&nats.StreamConfig{
		Name:     module,
		Subjects: []string{module + ".>"},
		Storage:  nats.FileStorage,
		MaxAge:   time.Hour * 24 * 7, //keep for a week
	}); err != nil {
		panic("Error in creating stream")
	}
}

func createConsumerIfMissing(sender, receiver string) {
	if _, err := JetStream.StreamInfo(sender); err != nil {
		panic("No stream `" + sender + "`")
	}

	if _, err := JetStream.ConsumerInfo(sender, receiver); err != nil {
		return
	}

	if _, err := JetStream.AddConsumer(sender, &nats.ConsumerConfig{
		Durable:       receiver,
		FilterSubject: sender + ".*",
	}); err != nil {
		panic("Error in creating stream")
	}
}
