package gamemanager

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManifestItem_UnmarshalJSON(t *testing.T) {
	jsonData := `{
		"FormatVersion": 1,
		"DisplayName": "Test Game",
		"InstallSize": 123456789,
		"MainWindowProcessName": "TestGame.exe",
		"ProcessNames": ["TestGame.exe", "Helper.exe"]
	}`

	var item ManifestItem
	err := json.Unmarshal([]byte(jsonData), &item)
	assert.NoError(t, err)
	assert.Equal(t, 1, item.FormatVersion)
	assert.Equal(t, "Test Game", item.DisplayName)
	assert.Equal(t, int64(123456789), item.InstallSize)
	assert.Equal(t, "TestGame.exe", item.MainWindowProcessName)
	assert.Equal(t, []string{"TestGame.exe", "Helper.exe"}, item.ProcessNames)
}

func TestLauncherInstalled_UnmarshalJSON(t *testing.T) {
	jsonData := `{
		"InstallLocation": "/games/test",
		"NamespaceId": "test-namespace",
		"ItemId": "test-item",
		"ArtifactId": "test-artifact",
		"AppVersion": "1.0.0",
		"AppName": "TestGame"
	}`

	var installed LauncherInstalled
	err := json.Unmarshal([]byte(jsonData), &installed)
	assert.NoError(t, err)
	assert.Equal(t, "/games/test", installed.InstallLocation)
	assert.Equal(t, "test-namespace", installed.NamespaceId)
	assert.Equal(t, "test-item", installed.ItemId)
	assert.Equal(t, "test-artifact", installed.ArtifactId)
	assert.Equal(t, "1.0.0", installed.AppVersion)
	assert.Equal(t, "TestGame", installed.AppName)
}

func TestEpicUrl_Format(t *testing.T) {
	manifest := &ManifestItem{
		CatalogNamespace: "namespace",
		CatalogItemId:    "itemid",
		AppName:          "appid",
	}
	expected := "com.epicgames.launcher://apps/namespace:itemid:appid?action=launch&silent=true"
	actual := FormatEpicUrl(manifest)
	assert.Equal(t, expected, actual)
}
