package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viper *viper.Viper) *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(viper.GetInt32("log.level")))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
