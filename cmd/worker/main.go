package main

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"golang-clean-architecture/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)

	consumerUser := config.NewKafkaConsumer(viperConfig, logger)
	signalUser := make(chan string)

	consumerContact := config.NewKafkaConsumer(viperConfig, logger)
	signalContact := make(chan string)

	consumerAddress := config.NewKafkaConsumer(viperConfig, logger)
	signalAddress := make(chan string)

	go config.ConsumeTopic(signalUser, consumerUser, "users", logger, func(message *kafka.Message) error {
		event := string(message.Value)
		logger.Infof("Received topic users with event: %s from partition %d", event, message.TopicPartition.Partition)
		return nil
	})

	go config.ConsumeTopic(signalContact, consumerContact, "contacts", logger, func(message *kafka.Message) error {
		event := string(message.Value)
		logger.Infof("Received topic contacts with event: %s from partition %d", event, message.TopicPartition.Partition)
		return nil
	})

	go config.ConsumeTopic(signalAddress, consumerAddress, "addresses", logger, func(message *kafka.Message) error {
		event := string(message.Value)
		logger.Infof("Received topic addresses with event: %s from partition %d", event, message.TopicPartition.Partition)
		return nil
	})

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false

	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down server gracefully, SIGNAL NAME :", s)
			signalUser <- "stop"
			signalContact <- "stop"
			signalAddress <- "stop"
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing

	//	./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic users --group golang-clean-architecture
	//$ ./bin/kafka-consumer-groups.sh --list --bootstrap-server localhost:9092
	// ./bin/kafka-consumer-groups.sh --describe --group golang-clean-architecture --members --bootstrap-server localhost:9092
}
