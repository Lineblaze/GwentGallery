package config

import (
	"github.com/Lineblaze/GwentGallery/app/internal/infrastructure/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

var Config struct {
	App struct {
		Jwt struct {
			AccessExpirationMinutes int16  `yaml:"access_expiration_minutes"`
			RefreshExpirationDays   int16  `yaml:"refresh_expiration_days"`
			AccessTokenSecretKey    string `yaml:"access_token_secret_key"`
			RefreshTokenSecretKey   string `yaml:"refresh_token_secret_key"`
		} `yaml:"jwt"`
	} `yaml:"app"`
	Interfaces struct {
		Http struct {
			Port uint32 `yaml:"port"`
			Host string `yaml:"host"`
		} `yaml:"http"`
	} `yaml:"interfaces"`
	Logger             logger.Config `yaml:"logger"`
	DatabaseConnString string        `yaml:"database_conn_string"`
}

func Init(configPath string) error {
	err := cleanenv.ReadConfig(configPath, &Config)
	return err
}
