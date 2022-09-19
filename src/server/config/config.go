package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Mongo struct {
	User     string `env:"mongo_user"`
	Pass     string `env:"mongo_pass"`
	Host     string `env:"mongo_host"`
	Args     string `env:"mongo_args"`
	Database string `env:"mongo_database"`
}

type Config struct {
	Environment string
	Version     string
	Mongo       Mongo
}

var Instance *Config = nil

func ReadConfigFromEnv(environment, version string) (Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return Config{}, fmt.Errorf(`error reading env: %w`, err)
	}

	cfg.Environment = environment
	cfg.Version = version

	Instance = &cfg

	return cfg, nil
}
