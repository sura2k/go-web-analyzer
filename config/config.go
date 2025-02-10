package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// init() will be triggered by GO when the package is loaded
func init() {
	// Get the config path from the environment variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable is not set")
	}

	viper.SetConfigName("config") // Config file name
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath) // Config path
	viper.AutomaticEnv()            // Advice to read from Env Variables as well. Note: Env has the higher precedency over config file

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("config file error: %s", err) // Note: Log.Fatalf() will execute os.Exit(1) to exit fom the application
	}

	// Unmarshal the config file properties into the global Config struct
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to map config properties into Config struct, %s", err)
	}

	log.Println("Configurations loaded")
}

// Global variable to hold the config values
var Config struct {
	Defaults struct {
		HTTP struct {
			Timeout struct {
				Seconds int `mapstructure:"seconds"`
			} `mapstructure:"timeout"`
		} `mapstructure:"http"`
	} `mapstructure:"defaults"`
	Analyzers struct {
		LinksAnalyzer struct {
			LinkHealthCheck struct {
				BatchSize int `mapstructure:"batch-size"`
			} `mapstructure:"link-health-check"`
		} `mapstructure:"LinksAnalyzer"`
	} `mapstructure:"analyzers"`
}
