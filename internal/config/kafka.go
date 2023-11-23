package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

func NewKafkaConsumer(config *viper.Viper) (*kafka.Consumer, error) {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.GetString("kafka.bootstrap.servers"),
		"group.id":          config.GetString("kafka.group.id"),
		"auto.offset.reset": config.GetString("kafka.auto.offset.reset"),
	}

	return kafka.NewConsumer(kafkaConfig)
}

func NewKafkaProducer(config *viper.Viper) (*kafka.Producer, error) {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.GetString("kafka.bootstrap.servers"),
	}

	return kafka.NewProducer(kafkaConfig)
}
