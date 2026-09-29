package config

import (
	"os"

	"github.com/cristianperen/servercheck/internal/models"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Servers []models.Server `yaml:"servers"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
