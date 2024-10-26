package types

import "time"

const (
	ModLoaderTypeAny        ModLoaderType = 0
	ModLoaderTypeForge      ModLoaderType = 1
	ModLoaderTypeCauldron   ModLoaderType = 2
	ModLoaderTypeLiteLoader ModLoaderType = 3
	ModLoaderTypeFabric     ModLoaderType = 4
	ModLoaderTypeQuilt      ModLoaderType = 5
	ModLoaderTypeNeoForge   ModLoaderType = 6
)

const (
	ModStatusNew             ModStatus = 1
	ModStatusChangesRequired ModStatus = 2
	ModStatusUnderSoftReview ModStatus = 3
	ModStatusApproved        ModStatus = 4
	ModStatusRejected        ModStatus = 5
	ModStatusChangesMade     ModStatus = 6
	ModStatusInactive        ModStatus = 7
	ModStatusAbandoned       ModStatus = 8
	ModStatusDeleted         ModStatus = 9
	ModStatusUnderReview     ModStatus = 10
)

type ModLoaderType int

type ModStatus int

type GameVersionsResponse struct {
	CurseforgeAPIResponse
	RawResponse
	Data GameVersions `json:"data"`
}

type GameVersions []GameVersion

type GameVersion struct {
	Id   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type MinecraftVersionsResponse struct {
	CurseforgeAPIResponse
	RawResponse
	Data MinecraftVersions `json:"data"`
}

type MinecraftVersions []MinecraftVersion

type MinecraftVersion struct {
	Id                    int       `json:"id"`
	GameVersionId         int       `json:"gameVersionId"`
	VersionString         string    `json:"versionString"`
	JarDownloadUrl        string    `json:"jarDownloadUrl"`
	JsonDownloadUrl       string    `json:"jsonDownloadUrl"`
	Approved              bool      `json:"approved"`
	DateModified          time.Time `json:"dateModified"`
	GameVersionTypeId     int       `json:"gameVersionTypeId"`
	GameVersionStatus     int       `json:"gameVersionStatus"`
	GameVersionTypeStatus int       `json:"gameVersionTypeStatus"`
}

type MinecraftModLoadersResponse struct {
	CurseforgeAPIResponse
	RawResponse
	Data MinecraftModLoaders `json:"data"`
}

type MinecraftModLoaders []MinecraftModLoader

type MinecraftModLoader struct {
	Name         string    `json:"name"`
	GameVersion  string    `json:"gameVersion"`
	Latest       bool      `json:"latest"`
	Recommended  bool      `json:"recommended"`
	DateModified time.Time `json:"dateModified"`
	Type         int       `json:"type"`
}
