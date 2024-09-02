package config

import (
	"fmt"
	"github.com/vrischmann/envconfig"
)

type Kafka struct {
	Brokers string
	Topic   string
	GroupID string
}

type Email struct {
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

type Config struct {
	Kafka
	Email
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Init(&cfg); err != nil {
		return Config{}, fmt.Errorf("get configs: %w", err)
	}

	return cfg, nil
}
