package gamemanager

// ManifestItem reflects the structure of ManifestExample.item
type ManifestItem struct {
	FormatVersion               int      `json:"FormatVersion"`
	BIsIncompleteInstall        bool     `json:"bIsIncompleteInstall"`
	LaunchCommand               string   `json:"LaunchCommand"`
	LaunchExecutable            string   `json:"LaunchExecutable"`
	ManifestLocation            string   `json:"ManifestLocation"`
	ManifestHash                string   `json:"ManifestHash"`
	BIsApplication              bool     `json:"bIsApplication"`
	BIsExecutable               bool     `json:"bIsExecutable"`
	BIsManaged                  bool     `json:"bIsManaged"`
	BNeedsValidation            bool     `json:"bNeedsValidation"`
	BRequiresAuth               bool     `json:"bRequiresAuth"`
	BAllowMultipleInstances     bool     `json:"bAllowMultipleInstances"`
	BCanRunOffline              bool     `json:"bCanRunOffline"`
	BAllowUriCmdArgs            bool     `json:"bAllowUriCmdArgs"`
	BLaunchElevated             bool     `json:"bLaunchElevated"`
	BaseURLs                    []string `json:"BaseURLs"`
	BuildLabel                  string   `json:"BuildLabel"`
	AppCategories               []string `json:"AppCategories"`
	ChunkDbs                    []string `json:"ChunkDbs"`
	CompatibleApps              []string `json:"CompatibleApps"`
	DisplayName                 string   `json:"DisplayName"`
	InstallationGuid            string   `json:"InstallationGuid"`
	InstallLocation             string   `json:"InstallLocation"`
	InstallSessionId            string   `json:"InstallSessionId"`
	InstallTags                 []string `json:"InstallTags"`
	InstallComponents           []string `json:"InstallComponents"`
	HostInstallationGuid        string   `json:"HostInstallationGuid"`
	PrereqIds                   []string `json:"PrereqIds"`
	PrereqSHA1Hash              string   `json:"PrereqSHA1Hash"`
	LastPrereqSucceededSHA1Hash string   `json:"LastPrereqSucceededSHA1Hash"`
	StagingLocation             string   `json:"StagingLocation"`
	TechnicalType               string   `json:"TechnicalType"`
	VaultThumbnailUrl           string   `json:"VaultThumbnailUrl"`
	VaultTitleText              string   `json:"VaultTitleText"`
	InstallSize                 int64    `json:"InstallSize"`
	MainWindowProcessName       string   `json:"MainWindowProcessName"`
	ProcessNames                []string `json:"ProcessNames"`
	BackgroundProcessNames      []string `json:"BackgroundProcessNames"`
	IgnoredProcessNames         []string `json:"IgnoredProcessNames"`
	DlcProcessNames             []string `json:"DlcProcessNames"`
	MandatoryAppFolderName      string   `json:"MandatoryAppFolderName"`
	OwnershipToken              string   `json:"OwnershipToken"`
	SidecarConfigRevision       int      `json:"SidecarConfigRevision"`
	OwnershipTokenForNs         string   `json:"OwnershipTokenForNs"`
	CatalogNamespace            string   `json:"CatalogNamespace"`
	CatalogItemId               string   `json:"CatalogItemId"`
	AppName                     string   `json:"AppName"`
	AppVersionString            string   `json:"AppVersionString"`
	MainGameCatalogNamespace    string   `json:"MainGameCatalogNamespace"`
	MainGameCatalogItemId       string   `json:"MainGameCatalogItemId"`
	MainGameAppName             string   `json:"MainGameAppName"`
	AllowedUriEnvVars           []string `json:"AllowedUriEnvVars"`
}

type LauncherInstalled struct {
	InstallLocation string `json:"InstallLocation"`
	NamespaceId     string `json:"NamespaceId"`
	ItemId          string `json:"ItemId"`
	ArtifactId      string `json:"ArtifactId"`
	AppVersion      string `json:"AppVersion"`
	AppName         string `json:"AppName"`
}

// EpicUrl is the URL schema to launch Epic Games Store games
// Example: com.epicgames.launcher://apps/NamespaceId:CatalogItemId:AppId?action=launch&silent=true
const EpicUrl = "com.epicgames.launcher://apps/%s:%s:%s?action=launch&silent=true"
