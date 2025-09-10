package gamemanager

import (
	"testing"

	"github.com/feed3r/play-harbor/go-launcher/config"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestNewGameManager(t *testing.T) {
	cfg := &config.Config{}
	fs := afero.NewMemMapFs()
	gm := NewGameManager(cfg, fs)
	assert.NotNil(t, gm)
	assert.Equal(t, cfg, gm.Config)
	assert.Equal(t, fs, gm.Fs)
}

func TestLoadLauncherInstalled_EmptyDir(t *testing.T) {
	cfg := &config.Config{}
	cfg.EpicGamesStore.LauncherInstalledPath = "/not-exist-dir"
	fs := afero.NewMemMapFs()
	gm := NewGameManager(cfg, fs)
	items, err := gm.LoadLauncherInstalled()
	assert.Error(t, err)
	assert.Nil(t, items)
}

func TestLoadManifestFile_NotFound(t *testing.T) {
	cfg := &config.Config{}
	fs := afero.NewMemMapFs()
	gm := NewGameManager(cfg, fs)
	_, err := gm.LoadManifestFile("/tmp/notfound.json")
	assert.Error(t, err)
}

func TestFillGameDescriptors_Empty(t *testing.T) {
	cfg := &config.Config{}
	cfg.EpicGamesStore.ManifestsFolderPath = "/not-exist-dir"
	fs := afero.NewMemMapFs()
	gm := NewGameManager(cfg, fs)
	err := gm.FillGameDescriptors()
	assert.Error(t, err)
}

func TestLoadLauncherInstalled_ValidFile(t *testing.T) {
	fs := afero.NewMemMapFs()
	path := "/LauncherInstalled.dat"
	content := `{"InstallationList": [{"NamespaceId": "test-ns", "InstallLocation": "/games/test", "AppName": "TestApp"}]}`
	err := afero.WriteFile(fs, path, []byte(content), 0644)
	assert.NoError(t, err)

	cfg := &config.Config{}
	cfg.EpicGamesStore.LauncherInstalledPath = path
	gm := NewGameManager(cfg, fs)
	items, err := gm.LoadLauncherInstalled()
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Contains(t, items, "test-ns")
	assert.Equal(t, "/games/test", items["test-ns"].InstallLocation)
	assert.Equal(t, "TestApp", items["test-ns"].AppName)
}

func TestLoadManifestFile_ValidFile(t *testing.T) {
	fs := afero.NewMemMapFs()
	path := "/manifest.json"
	content := `{"DisplayName": "Test Game", "AppName": "TestApp", "CatalogNamespace": "ns1", "InstallLocation": "/games/test"}`
	err := afero.WriteFile(fs, path, []byte(content), 0644)
	assert.NoError(t, err)

	cfg := &config.Config{}
	gm := NewGameManager(cfg, fs)
	item, err := gm.LoadManifestFile(path)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, "Test Game", item.DisplayName)
	assert.Equal(t, "TestApp", item.AppName)
	assert.Equal(t, "ns1", item.CatalogNamespace)
	assert.Equal(t, "/games/test", item.InstallLocation)
}

func TestFillGameDescriptors_PopulatesGames(t *testing.T) {
	fs := afero.NewMemMapFs()
	manifestDir := "/manifests"
	_ = fs.Mkdir(manifestDir, 0755)
	manifestPath := manifestDir + "/game1.json"
	manifestContent := `{"DisplayName": "Test Game", "AppName": "TestApp", "CatalogNamespace": "ns1", "InstallLocation": "/games/test"}`
	err := afero.WriteFile(fs, manifestPath, []byte(manifestContent), 0644)
	assert.NoError(t, err)

	launcherPath := "/LauncherInstalled.dat"
	launcherContent := `{"InstallationList": [{"NamespaceId": "ns1", "InstallLocation": "/games/test", "AppName": "TestApp"}]}`
	err = afero.WriteFile(fs, launcherPath, []byte(launcherContent), 0644)
	assert.NoError(t, err)

	cfg := &config.Config{}
	cfg.EpicGamesStore.ManifestsFolderPath = manifestDir
	cfg.EpicGamesStore.LauncherInstalledPath = launcherPath
	gm := NewGameManager(cfg, fs)
	err = gm.FillGameDescriptors()
	assert.NoError(t, err)
	assert.NotNil(t, gm.Games)
	assert.Len(t, gm.Games, 1)
	assert.Equal(t, "Test Game", gm.Games[0].DisplayName)
}

func TestFillGameDescriptors_ManifestWithoutLauncherInstalled(t *testing.T) {
	fs := afero.NewMemMapFs()
	manifestDir := "/manifests"
	_ = fs.Mkdir(manifestDir, 0755)
	manifestPath := manifestDir + "/game2.json"
	manifestContent := `{"DisplayName": "Orphan Game", "AppName": "OrphanApp", "CatalogNamespace": "missing-ns", "InstallLocation": "/games/orphan"}`
	err := afero.WriteFile(fs, manifestPath, []byte(manifestContent), 0644)
	assert.NoError(t, err)

	launcherPath := "/LauncherInstalled.dat"
	launcherContent := `{"InstallationList": [{"NamespaceId": "ns1", "InstallLocation": "/games/test", "AppName": "TestApp"}]}`
	err = afero.WriteFile(fs, launcherPath, []byte(launcherContent), 0644)
	assert.NoError(t, err)

	cfg := &config.Config{}
	cfg.EpicGamesStore.ManifestsFolderPath = manifestDir
	cfg.EpicGamesStore.LauncherInstalledPath = launcherPath
	gm := NewGameManager(cfg, fs)
	err = gm.FillGameDescriptors()
	assert.NoError(t, err)
	assert.Len(t, gm.Games, 0)
}

func TestFillGameDescriptors_EndToEnd(t *testing.T) {
	fs := afero.NewMemMapFs()
	manifestDir := "/manifests"
	_ = fs.Mkdir(manifestDir, 0755)
	manifestPath := manifestDir + "/game1.json"
	manifestContent := `{"DisplayName": "End2End Game", "AppName": "End2EndApp", "CatalogNamespace": "ns-end2end", "InstallLocation": "/games/end2end", "LaunchExecutable": "end2end.exe", "CatalogItemId": "item123"}`
	err := afero.WriteFile(fs, manifestPath, []byte(manifestContent), 0644)
	assert.NoError(t, err)

	launcherPath := "/LauncherInstalled.dat"
	launcherContent := `{"InstallationList": [{"NamespaceId": "ns-end2end", "InstallLocation": "/games/end2end", "AppName": "End2EndApp"}]}`
	err = afero.WriteFile(fs, launcherPath, []byte(launcherContent), 0644)
	assert.NoError(t, err)

	cfg := &config.Config{}
	cfg.EpicGamesStore.ManifestsFolderPath = manifestDir
	cfg.EpicGamesStore.LauncherInstalledPath = launcherPath
	gm := NewGameManager(cfg, fs)
	err = gm.FillGameDescriptors()
	assert.NoError(t, err)
	assert.NotNil(t, gm.Games)
	assert.Len(t, gm.Games, 1)
	game := gm.Games[0]
	assert.Equal(t, "End2End Game", game.DisplayName)
	assert.Equal(t, "end2end.exe", game.ExeName)
	assert.Contains(t, game.EpicUrl, "item123") // EpicUrl dovrebbe contenere l'itemId
}
