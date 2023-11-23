package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"golang-clean-architecture/internal/config"
	"strconv"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)

	producerUser := config.NewKafkaProducer(viperConfig, logger)

	topic := "users"
	for i := 0; i < 10; i++ {
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("Hello World " + strconv.Itoa(i)),
		}
		err := producerUser.Produce(message, nil)
		logger.Info("Producing message: ", message)
		if err != nil {
			logger.Fatalf("Failed to produce message: %v", err)
		}
	}

	time.Sleep(5 * time.Second) // wait for all messages to be delivered
	producerUser.Close()
}
