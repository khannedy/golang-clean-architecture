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

	webPort := viperConfig.GetInt("web.port")
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	//start server
	err = app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatal(err)
	}
}
