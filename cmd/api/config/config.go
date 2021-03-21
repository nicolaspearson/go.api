package config

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Config is a global object that holds all runtime variables.
var Config config

type config struct {
	Database *gorm.DB
}

// Vars is a global object that holds all configuration / environment variables.
var Vars vars

type vars struct {
	DbName         string
	DbUser         string
	DbPassword     string
	DbHost         string
	DbPort         string
	Environment    string
	ReleaseVersion string
	ServerHost     string
	ServerPort     string
	Version        string
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
	v.BindEnv("releaseVersion", "API_RELEASE_VERSION")
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read the configuration file: %v", err)
	}
	return v.Unmarshal(&Vars)
}
