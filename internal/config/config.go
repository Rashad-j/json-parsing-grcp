package config

import "github.com/caarlos0/env"

type Config struct {
	Addr           string `env:"ADDR" envDefault:":8080"`
	FilesDirectory string `env:"FILES_DIR" envDefault:"./files"`
	Extension      string `env:"EXTENSION"  envDefault:"json"`
}

func ReadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
