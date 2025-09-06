package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

const TEST_CONFIG_YAML = `EpicGamesStore:
    executable: "EpicGamesLauncher.exe"
Global:
    sleep_with_manager: 10s
    sleep_without_manager: 30s
    max_polling_attempts: 20
    polling_interval: 1s
`

func TestLoadConfig(t *testing.T) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(TEST_CONFIG_YAML), &cfg); err != nil {
		t.Fatalf("yaml.Unmarshal failed: %v", err)
	}

	if cfg.EpicGamesStore.Executable != "EpicGamesLauncher.exe" {
		t.Errorf("EpicGamesStore.Executable: got %q, want %q", cfg.EpicGamesStore.Executable, "EpicGamesLauncher.exe")
	}
	if cfg.Global.SleepWithManager != 10_000_000_000 {
		t.Errorf("Global.SleepWithManager: got %v, want 10s", cfg.Global.SleepWithManager)
	}
	if cfg.Global.MaxPollingAttempts != 20 {
		t.Errorf("Global.MaxPollingAttempts: got %d, want 20", cfg.Global.MaxPollingAttempts)
	}
}
