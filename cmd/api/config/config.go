package config

import (
	"log"

	"github.com/spf13/viper"
)

// Vars is a global object that holds all configuration / environment variables.
var Vars vars

type vars struct {
	DbName              string
	DbUser              string
	DbPassword          string
	DbHost              string
	DbPort              string
	Environment         string
	RabbitMqUsername    string
	RabbitMqPassword    string
	RabbitMqHost        string
	RabbitMqVirtualHost string
	ReleaseVersion      string
	ServerHost          string
	ServerPort          string
	Version             string
}

// LoadConfig loads config variables from file paths
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("api")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("api")
	v.AutomaticEnv()
	v.BindEnv("dbName", "API_DB_NAME")
	v.BindEnv("dbUser", "API_DB_USER")
	v.BindEnv("dbPassword", "API_DB_PASSWORD")
	v.BindEnv("dbHost", "API_DB_HOST")
	v.BindEnv("dbPort", "API_DB_PORT")
	v.BindEnv("rabbitMqUsername", "API_RABBIT_MQ_USERNAME")
	v.BindEnv("rabbitMqPassword", "API_RABBIT_MQ_PASSWORD")
	v.BindEnv("rabbitMqHost", "API_RABBIT_MQ_HOST")
	v.BindEnv("rabbitMqVirtualHost", "API_RABBIT_MQ_VIRTUAL_HOST")
	v.BindEnv("releaseVersion", "API_RELEASE_VERSION")
	v.BindEnv("serverHost", "API_SERVER_HOST")
	v.BindEnv("serverPort", "API_SERVER_PORT")
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read the configuration file: %v", err)
	}
	return v.Unmarshal(&Vars)
}
