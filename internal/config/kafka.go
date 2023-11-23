package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func NewKafkaConsumer(config *viper.Viper, log *logrus.Logger) *kafka.Consumer {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.GetString("kafka.bootstrap.servers"),
		"group.id":          config.GetString("kafka.group.id"),
		"auto.offset.reset": config.GetString("kafka.auto.offset.reset"),
	}

	consumer, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	return consumer
}

func ConsumeTopic(signal chan string, consumer *kafka.Consumer, topic string, log *logrus.Logger, handler func(message *kafka.Message) error) {
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

func NewKafkaProducer(config *viper.Viper, log *logrus.Logger) *kafka.Producer {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.GetString("kafka.bootstrap.servers"),
	}

	producer, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	return producer
}
