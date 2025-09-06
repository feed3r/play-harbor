package config

import (
	"bytes"
	"io"
	"testing"
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
	ReadConfigFile = func(path string) (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewBufferString(TEST_CONFIG_YAML)), nil
	}

	cfg, err := LoadConfig("dummy_path.yaml")
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
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
