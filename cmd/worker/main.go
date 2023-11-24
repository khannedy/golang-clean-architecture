package main

import (
	"golang-clean-architecture/internal/config"
	"golang-clean-architecture/internal/delivery/messaging"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting worker service")

	logger.Info("setup user consumer")
	userConsumer := config.NewKafkaConsumer(viperConfig, logger)
	userSignal := make(chan string)
	userHandler := messaging.NewUserConsumer(logger)
	go config.ConsumeTopic(userSignal, userConsumer, "users", logger, userHandler.Consume)

	logger.Info("setup contact consumer")
	contactConsumer := config.NewKafkaConsumer(viperConfig, logger)
	contactSignal := make(chan string)
	contactHandler := messaging.NewContactConsumer(logger)
	go config.ConsumeTopic(contactSignal, contactConsumer, "contacts", logger, contactHandler.Consume)

	logger.Info("setup address consumer")
	addressConsumer := config.NewKafkaConsumer(viperConfig, logger)
	addressSignal := make(chan string)
	addressHandler := messaging.NewAddressConsumer(logger)
	go config.ConsumeTopic(addressSignal, addressConsumer, "addresses", logger, addressHandler.Consume)

	logger.Info("Worker is running")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down server gracefully, SIGNAL NAME :", s)
			userSignal <- "stop"
			contactSignal <- "stop"
			addressSignal <- "stop"
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}
