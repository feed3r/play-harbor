package config

import (
	"io"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type EpicGamesStoreConfig struct {
	Executable            string `yaml:"executable"`
	LauncherInstalledPath string `yaml:"launcher_installed_path,omitempty"`
	ManifestsFolderPath   string `yaml:"manifests_path,omitempty"`
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

// Global mockable function to read the configuration file
var ReadConfigFile = func(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func LoadConfig(path string) (*Config, error) {
	r, err := ReadConfigFile(path)
	if err != nil {
		// File not found: return default config
		return &Config{
			Global: GlobalConfig{
				SleepWithManager:    30 * time.Second,
				SleepWithoutManager: 10 * time.Second,
				MaxPollingAttempts:  10,
				PollingInterval:     1 * time.Second,
			},
			EpicGamesStore: EpicGamesStoreConfig{
				Executable:            "EpicGamesLauncher.exe",
				LauncherInstalledPath: "c:\\ProgramData\\Epic\\UnrealEngineLauncher\\LauncherInstalled.dat",
				ManifestsFolderPath:   "c:\\ProgramData\\Epic\\EpicGamesLauncher\\Data\\Manifests",
			},
		}, nil
	}
	defer r.Close()

	var cfg Config
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
