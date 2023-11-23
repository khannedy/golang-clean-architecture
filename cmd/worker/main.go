package main

import (
	"fmt"
	"golang-clean-architecture/internal/config"
)

func main() {
	viperConfig, err := config.NewViper()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	consumer, err := config.NewKafkaConsumer(viperConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error kafka consumer: %w \n", err))
	}

	err = consumer.Close()
	if err != nil {
		panic(fmt.Errorf("Fatal error kafka consumer close: %w \n", err))
	}
}
