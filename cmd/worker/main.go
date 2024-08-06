package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	go RunUserConsumer(logger, viperConfig, ctx)
	go RunContactConsumer(logger, viperConfig, ctx)
	go RunAddressConsumer(logger, viperConfig, ctx)

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

func RunAddressConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup address consumer")
	addressConsumer := config.NewKafkaConsumer(viperConfig, logger)
	addressHandler := messaging.NewAddressConsumer(logger)
	messaging.ConsumeTopic(ctx, addressConsumer, "addresses", logger, addressHandler.Consume)
}

func RunContactConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup contact consumer")
	contactConsumer := config.NewKafkaConsumer(viperConfig, logger)
	contactHandler := messaging.NewContactConsumer(logger)
	messaging.ConsumeTopic(ctx, contactConsumer, "contacts", logger, contactHandler.Consume)
}

func RunUserConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup user consumer")
	userConsumer := config.NewKafkaConsumer(viperConfig, logger)
	userHandler := messaging.NewUserConsumer(logger)
	messaging.ConsumeTopic(ctx, userConsumer, "users", logger, userHandler.Consume)
}
