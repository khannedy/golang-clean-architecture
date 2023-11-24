package messaging

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"time"
)

type ConsumerHandler func(message *kafka.Message) error

func ConsumeTopic(ctx context.Context, consumer *kafka.Consumer, topic string, log *logrus.Logger, handler ConsumerHandler) {
	err := consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			message, err := consumer.ReadMessage(time.Second)
			if err == nil {
				err := handler(message)
				if err != nil {
					log.Errorf("Failed to process message: %v", err)
				} else {
					_, err = consumer.CommitMessage(message)
					if err != nil {
						log.Fatalf("Failed to commit message: %v", err)
					}
				}
			} else if !err.(kafka.Error).IsTimeout() {
				log.Warnf("Consumer error: %v (%v)\n", err, message)
			}
		}
	}

	log.Infof("Closing consumer for topic : %s", topic)
	err = consumer.Close()
	if err != nil {
		panic(err)
	}
}
