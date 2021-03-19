package config

import (
	"log"
	"github.com/spf13/viper"
)

// const (
// 	bankService  = "BANK_SERVICE"
// 	fraudService = "FRAUD_SERVICE"
// )

type Config struct {
	httpServerConfig HTTPServerConfig
}

func (config Config) GetHTTPServerConfig() HTTPServerConfig {
    log.Print("GetHTTPServerConfig", config)
	return config.httpServerConfig
}

func NewConfig(configFile string) Config {
	viper.AutomaticEnv()

	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	return Config{
		httpServerConfig: newHTTPServerConfig(),
	}
}
