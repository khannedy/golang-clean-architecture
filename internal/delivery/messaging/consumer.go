package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"time"
)

type ConsumerHandler func(message *kafka.Message) error

func ConsumeTopic(signal chan string, consumer *kafka.Consumer, topic string, log *logrus.Logger, handler ConsumerHandler) {
	err := consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	stop := false

	for !stop {
		select {
		case <-signal:
			log.Info("Got one of stop signals, shutting down server gracefully")
			stop = true
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

	log.Info("Closing consumer")
	err = consumer.Close()
	if err != nil {
		panic(err)
	}
}
