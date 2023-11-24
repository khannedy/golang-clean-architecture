package main

import (
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("setup user consumer")
	userConsumer := config.NewKafkaConsumer(viperConfig, logger)
	userHandler := messaging.NewUserConsumer(logger)
	go messaging.ConsumeTopic(ctx, userConsumer, "users", logger, userHandler.Consume)

	logger.Info("setup contact consumer")
	contactConsumer := config.NewKafkaConsumer(viperConfig, logger)
	contactHandler := messaging.NewContactConsumer(logger)
	go messaging.ConsumeTopic(ctx, contactConsumer, "contacts", logger, contactHandler.Consume)

	logger.Info("setup address consumer")
	addressConsumer := config.NewKafkaConsumer(viperConfig, logger)
	addressHandler := messaging.NewAddressConsumer(logger)
	go messaging.ConsumeTopic(ctx, addressConsumer, "addresses", logger, addressHandler.Consume)

	logger.Info("Worker is running")

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}
