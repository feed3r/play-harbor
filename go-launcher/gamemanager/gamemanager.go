package gamemanager

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/feed3r/play-harbor/go-launcher/config"
)

type GameDescriptor struct {
	DisplayName string
	EpicUrl     string
	ExeName     string
	Pid         int
}

type GameManager struct {
	Config *config.Config
	Games  []*GameDescriptor
}

func NewGameManager(cfg *config.Config) *GameManager {
	return &GameManager{
		Config: cfg,
	}
}

// LoadGames loads LauncherInstalled.dat, parses InstallationList,
// and returns a map of LauncherInstalled indexed by NamespaceId.
func (r *GameManager) LoadLauncherInstalled() (map[string]*LauncherInstalled, error) {
	file, err := os.Open(r.Config.EpicGamesStore.LauncherInstalledPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Structure to match the file format
	var data struct {
		InstallationList []LauncherInstalled `json:"InstallationList"`
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	result := make(map[string]*LauncherInstalled)
	for _, item := range data.InstallationList {
		result[item.NamespaceId] = &item
	}
	return result, nil
}

func (r *GameManager) LoadManifestFile(manifestFilePath string) (*ManifestItem, error) {
	file, err := os.Open(manifestFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var item ManifestItem
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *GameManager) FillGameDescriptors() error {
	var gameDescriptors []*GameDescriptor
	r.Games = gameDescriptors

	//Fill the initial list of games
	launchers, err := r.LoadLauncherInstalled()
	if err != nil {
		return err
	}

	//Parse the manifests files
	manifestsFolder := r.Config.EpicGamesStore.ManifestsFolderPath
	files, err := os.ReadDir(manifestsFolder)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			manifest, err := r.LoadManifestFile(file.Name())
			if err != nil {
				return err
			}

			//Check that the manifest has a valid NameSpaceID
			if manifest.CatalogNamespace == "" || launchers[manifest.CatalogNamespace] == nil {
				//TODO: Write some log....
				continue
			}

			// Build the Epic URL
			launcherGameLink := fmt.Sprintf(EpicUrl, manifest.CatalogNamespace, manifest.CatalogItemId, manifest.AppName)

			r.Games = append(r.Games, &GameDescriptor{
				DisplayName: manifest.DisplayName,
				EpicUrl:     launcherGameLink,
				ExeName:     manifest.LaunchExecutable,
			})
		}

	}

	return nil
}
