package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type EpicGamesStoreConfig struct {
	Executable string `yaml:"executable"`
}

type Config struct {
	EpicGamesStore EpicGamesStoreConfig `yaml:"EpicGamesStore"`
	Global         GlobalConfig         `yaml:"Global"`
}

type GlobalConfig struct {
	SleepWithManager    time.Duration `yaml:"sleep_with_manager"`
	SleepWithoutManager time.Duration `yaml:"sleep_without_manager"`
	MaxPollingAttempts  int           `yaml:"max_polling_attempts"`
	PollingInterval     time.Duration `yaml:"polling_interval"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
