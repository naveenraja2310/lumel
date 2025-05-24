/*
settings/config.go
Author: Naveenraj O M
Modified By:
Description: This file contains the configuration settings for the application.
The configuration settings are stored in the Configuration struct and can be accessed throughout the application
*/

package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	Environment string `mapstructure:"environment"`
	Logger      LoggerConfig
	DBURI       string `mapstructure:"DB_URI"`
	DBName      string `mapstructure:"DB_NAME"`
	DB_TIME     int    `mapstructure:"DB_TIME"`
	AppPort     string `mapstructure:"APP_PORT"`
}

type LoggerConfig struct {
	FileName     string `mapstructure:"fileName"`
	FileSize     int    `mapstructure:"fileSize"`
	MaxLogFile   int    `mapstructure:"maxLogFile"`
	MaxRetention int    `mapstructure:"maxRetention"`
	CompressLog  bool   `mapstructure:"compressLog"`
	Level        string `mapstructure:"level"`
}

func InitConfig() (Configuration, error) {
	var configDir, envDir string

	currentDir, err := os.Getwd()
	if err != nil {
		return Config, err
	}

	configDir = filepath.Join(currentDir, "")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		configDir = "./"
	}

	envDir = filepath.Join(currentDir, "")
	if _, err := os.Stat(envDir); os.IsNotExist(err) {
		envDir = "./"
	}

	// Load `config.yaml`
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return Config, err
	}

	// Load `.env`
	viper.AddConfigPath(envDir)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.MergeInConfig(); err != nil {
		return Config, err
	}

	// Bind sensitive env variables
	envVars := []string{
		"DB_URI", "DB_NAME", "DB_TIME", "APP_PORT",
	}

	for _, envVar := range envVars {
		if err := viper.BindEnv(envVar); err != nil {
			return Config, err
		}
	}

	// Unmarshal the combined configuration
	if err := viper.Unmarshal(&Config); err != nil {
		return Config, err
	}

	if Config.Environment != "production" {
		Config.DB_TIME = 1000
	}

	if Config.DBURI == "" {
		log.Fatal("DB_URI is not set")
	}

	if Config.DBName == "" {
		log.Fatal("DB_NAME is not set")
	}

	return Config, nil
}
