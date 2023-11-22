package internal

import "github.com/spf13/viper"

// New is a function to load config from config.json
// You can change the implementation, for example load from env file, consul, etcd, etc
func New() (*viper.Viper, error) {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	err := config.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return config, nil
}
