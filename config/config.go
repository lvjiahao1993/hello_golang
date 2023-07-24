package config

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

// Config struct holds the configuration data
type Config struct {
	ServerPort  int    `mapstructure:"server_port"`
	DatabaseURL string `mapstructure:"database_url"`
	LogFilePath string `mapstructure:"log_file_path"`
	LogLevel    string `mapstructure:"log_level"`
}

// LoadConfig loads the configuration from file and environment variables
func LoadConfig() (*Config, error) {
	// Set default values for configuration
	viper.SetDefault("server_port", 8080)
	viper.SetDefault("database_url", "localhost:5432")
	viper.SetDefault("log_file_path", "app.log")
	viper.SetDefault("log_level", "info")

	// Read configuration from file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Search for the config file in the current directory

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Using default values.")
		} else {
			return nil, fmt.Errorf("failed to read configuration file: %s", err)
		}
	}

	// Read configuration from environment variables
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	// Unmarshal the configuration into the Config struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %s", err)
	}

	return &cfg, nil
}

// InitDB initializes the database connection
func InitDB(databaseURL string) (*sql.DB, error) {
	// Open a new database connection
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Check the database connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Database connection established")
	return db, nil
}
