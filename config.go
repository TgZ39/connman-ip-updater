package main

import (
	"errors"
	"github.com/pelletier/go-toml"
	"os"
)

const configFile = "ip_updater_config.toml"

type Config struct {
	Enabled             bool
	Domain              string
	WireguardConfigFile string
}

func GetConfig() (*Config, error) {
	buf, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(buf, &cfg)
	if err != nil {
		return nil, err
	}

	if len(cfg.Domain) == 0 || len(cfg.WireguardConfigFile) == 0 {
		return nil, errors.New("config has missing fields")
	}

	return &cfg, nil
}
