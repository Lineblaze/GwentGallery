package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	HTTP          struct {
		IP           string        `yaml:"ip" env:"HTTP-IP"`
		Port         int           `yaml:"port" env:"HTTP-PORT"`
		ReadTimeout  time.Duration `yaml:"read-timeout" env:"HTTP-READ-TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write-timeout" env:"HTTP-WRITE-TIMEOUT"`
		CORS         struct {
			AllowedMethods     []string `yaml:"allowed_methods" env:"HTTP-CORS-ALLOWED-METHODS"`
			AllowedOrigins     []string `yaml:"allowed_origins" env:"HTTP-CORS-ALLOWED-ORIGINS"`
			AllowCredentials   bool     `yaml:"allow_credentials" env:"HTTP-CORS-ALLOWED-CREDENTIALS"`
			AllowedHeaders     []string `yaml:"allowed_headers" env:"HTTP-CORS-ALLOWED-HEADERS"`
			OptionsPassthrough bool     `yaml:"options_passthrough" env:"HTTP-CORS-OPTIONS-PASSTHROUGH"`
			ExposedHeaders     []string `yaml:"exposed_headers" env:"HTTP-CORS-EXPOSED-HEADERS"`
			Debug              bool     `yaml:"debug" env:"HTTP-CORS-DEBUG"`
		} `yaml:"cors"`
	} `yaml:"http"`

	AppConfig struct {
		LogLevel  string `env:"LOG_LEVEL" env-default:"trace"`
		AdminUser struct {
			Email    string `env:"ADMIN-EMAIL" env-default:"admin"`
			Password string `env:"ADMIN-PASSWORD" env-default:"admin"`
		}
	} `yaml:"app"`

	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	}
}

const (
	EnvConfigPathName  = "CONFIG-PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		flag.StringVar(&configPath, FlagConfigPathName, "configs/config.local.yaml", "this is app config file")
		flag.Parse()

		log.Print("config init")

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal("config path is required")
		}

		instance = &Config{}

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			helpText := "GwentGallery"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
