package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
