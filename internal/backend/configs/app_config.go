package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	JwtSecret  string
}

var BackendConfig config

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

func GetDBConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		BackendConfig.DBUser,
		BackendConfig.DBPassword,
		BackendConfig.DBHost,
		BackendConfig.DBPort,
		BackendConfig.DBName,
	)
}

func GetJwtSecret() string {
	return BackendConfig.JwtSecret
}
