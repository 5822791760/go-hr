package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	JwtSecret  string
}

var AppConfig Config

func LoadConfig() error {
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	AppConfig = Config{
		DBHost:     viper.GetString("POSTGRES_HOST"),
		DBPort:     viper.GetString("POSTGRES_PORT"),
		DBName:     viper.GetString("POSTGRES_DB"),
		DBUser:     viper.GetString("POSTGRES_USER"),
		DBPassword: viper.GetString("POSTGRES_PASSWORD"),
		JwtSecret:  viper.GetString("JWT_SECRET"),
	}

	return nil
}

func GetDBConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
	)
}

func GetJwtSecret() string {
	return AppConfig.JwtSecret
}
