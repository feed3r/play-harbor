package config

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	require.NoError(t, err, "LoadConfig fallita")

	assert.Equal(t, "EpicGamesLauncher.exe", cfg.EpicGamesStore.Executable, "EpicGamesStore.Executable")
	assert.Equal(t, time.Duration(10_000_000_000), cfg.Global.SleepWithManager, "Global.SleepWithManager")
	assert.Equal(t, 20, cfg.Global.MaxPollingAttempts, "Global.MaxPollingAttempts")
}
