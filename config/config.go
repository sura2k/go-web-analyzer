package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// init() will be triggered by GO when the package is loaded
func init() {

	// Set default config values
	setDefaults()

	// Get the config path from the environment variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Println("WARNING: CONFIG_PATH environment variable is not set")
		configPath = "."
	}

	viper.SetConfigName("config") // Config file name
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath) // Config path
	viper.AutomaticEnv()            // Advice to read from Env Variables as well. Note: Env has the higher precedency over config file

	if err := viper.ReadInConfig(); err != nil {
		log.Println("WARNING: Config file not found. Using default values")
	} else {
		log.Println("Config file loaded successfully")
	}

	// Unmarshal the config file properties into the global Config struct
	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("unable to map config properties into Config struct, %s", err)
	}

	log.Println("Configurations loaded")
}

// Set default values
func setDefaults() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("defaults.http.timeout.seconds", 5)
	viper.SetDefault("analyzers.LinksAnalyzer.link-health-check.batch-size", 10)
}

// Global variable to hold the config values
var Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
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
