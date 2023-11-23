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

	log := config.NewLogger(viperConfig)
	log.Info("Start application")

	db, err := config.NewDatabase(viperConfig, log)
	if err != nil {
		panic(fmt.Errorf("Fatal error database: %w \n", err))
	}

	validate := config.NewValidator(viperConfig)

	app := config.NewFiber(viperConfig)

	producer, err := config.NewKafkaProducer(viperConfig)
	if err != nil {
		panic(fmt.Errorf("Fatal error kafka producer: %w \n", err))
	}

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
		Producer: producer,
	})

	//start server
	webPort := viperConfig.GetInt("web.port")
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatal(err)
	}
}
