package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config da aplicação
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	RabbitURL  string
	ServerPort int
}

// LoadConfig carrega as configs de variáveis de ambiente
func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	// Definindo padrões
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("SERVER_PORT", 8080)

	config := Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetInt("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		RabbitURL:  viper.GetString("RABBIT_URL"),
		ServerPort: viper.GetInt("SERVER_PORT"),
	}

	// Valida configs necessárias
	if config.DBHost == "" || config.DBUser == "" || config.DBPassword == "" || config.DBName == "" || config.RabbitURL == "" {
		return nil, fmt.Errorf("missing required configuration")
	}

	return &config, nil
}
