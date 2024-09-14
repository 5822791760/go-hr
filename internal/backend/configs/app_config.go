// Package configs provide config from .env to serve backend.
package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// All Avaliable config
type config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	JwtSecret  string
}

var BackendConfig config

// This will be call on started to load config into variable
func LoadConfig() error {
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	BackendConfig = config{
		DBHost:     viper.GetString("POSTGRES_HOST"),
		DBPort:     viper.GetString("POSTGRES_PORT"),
		DBName:     viper.GetString("POSTGRES_DB"),
		DBUser:     viper.GetString("POSTGRES_USER"),
		DBPassword: viper.GetString("POSTGRES_PASSWORD"),
		JwtSecret:  viper.GetString("JWT_SECRET"),
	}

	return nil
}

// For database connection
func GetDBConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		BackendConfig.DBUser,
		BackendConfig.DBPassword,
		BackendConfig.DBHost,
		BackendConfig.DBPort,
		BackendConfig.DBName,
	)
}

// For jwt token
func GetJwtSecret() string {
	return BackendConfig.JwtSecret
}
