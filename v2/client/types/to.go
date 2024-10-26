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

type ModsResponse struct {
	CurseforgeAPIResponse
	RawResponse
	Data       []ModData  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Links struct {
	WebsiteURL string `json:"websiteUrl"`
	WikiURL    string `json:"wikiUrl"`
	IssuesURL  string `json:"issuesUrl"`
	SourceURL  string `json:"sourceUrl"`
}
type Categories struct {
	ID               int       `json:"id"`
	GameID           int       `json:"gameId"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	URL              string    `json:"url"`
	IconURL          string    `json:"iconUrl"`
	DateModified     time.Time `json:"dateModified"`
	IsClass          bool      `json:"isClass"`
	ClassID          int       `json:"classId"`
	ParentCategoryID int       `json:"parentCategoryId"`
	DisplayIndex     int       `json:"displayIndex"`
}
type Authors struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Logo struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}
type Screenshots struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}
type Hashes struct {
	Value string `json:"value"`
	Algo  int    `json:"algo"`
}
type SortableGameVersions struct {
	GameVersionName        string    `json:"gameVersionName"`
	GameVersionPadded      string    `json:"gameVersionPadded"`
	GameVersion            string    `json:"gameVersion"`
	GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
	GameVersionTypeID      int       `json:"gameVersionTypeId"`
}
type Dependencies struct {
	ModID        int `json:"modId"`
	RelationType int `json:"relationType"`
}
type Modules struct {
	Name        string `json:"name"`
	Fingerprint int    `json:"fingerprint"`
}
type LatestFiles struct {
	ID                   int                    `json:"id"`
	GameID               int                    `json:"gameId"`
	ModID                int                    `json:"modId"`
	IsAvailable          bool                   `json:"isAvailable"`
	DisplayName          string                 `json:"displayName"`
	FileName             string                 `json:"fileName"`
	ReleaseType          int                    `json:"releaseType"`
	FileStatus           int                    `json:"fileStatus"`
	Hashes               []Hashes               `json:"hashes"`
	FileDate             time.Time              `json:"fileDate"`
	FileLength           int                    `json:"fileLength"`
	DownloadCount        int                    `json:"downloadCount"`
	FileSizeOnDisk       int                    `json:"fileSizeOnDisk"`
	DownloadURL          string                 `json:"downloadUrl"`
	GameVersions         []string               `json:"gameVersions"`
	SortableGameVersions []SortableGameVersions `json:"sortableGameVersions"`
	Dependencies         []Dependencies         `json:"dependencies"`
	ExposeAsAlternative  bool                   `json:"exposeAsAlternative"`
	ParentProjectFileID  int                    `json:"parentProjectFileId"`
	AlternateFileID      int                    `json:"alternateFileId"`
	IsServerPack         bool                   `json:"isServerPack"`
	ServerPackFileID     int                    `json:"serverPackFileId"`
	IsEarlyAccessContent bool                   `json:"isEarlyAccessContent"`
	EarlyAccessEndDate   time.Time              `json:"earlyAccessEndDate"`
	FileFingerprint      int                    `json:"fileFingerprint"`
	Modules              []Modules              `json:"modules"`
}
type LatestFilesIndexes struct {
	GameVersion       string        `json:"gameVersion"`
	FileID            int           `json:"fileId"`
	Filename          string        `json:"filename"`
	ReleaseType       int           `json:"releaseType"`
	GameVersionTypeID int           `json:"gameVersionTypeId"`
	ModLoader         ModLoaderType `json:"modLoader"`
}

type LatestEarlyAccessFilesIndexes struct {
	GameVersion       string `json:"gameVersion"`
	FileID            int    `json:"fileId"`
	Filename          string `json:"filename"`
	ReleaseType       int    `json:"releaseType"`
	GameVersionTypeID int    `json:"gameVersionTypeId"`
	ModLoader         int    `json:"modLoader"`
}

type ModData struct {
	ID                            int                             `json:"id"`
	GameID                        int                             `json:"gameId"`
	Name                          string                          `json:"name"`
	Slug                          string                          `json:"slug"`
	Links                         Links                           `json:"links"`
	Summary                       string                          `json:"summary"`
	Status                        int                             `json:"status"`
	DownloadCount                 int                             `json:"downloadCount"`
	IsFeatured                    bool                            `json:"isFeatured"`
	PrimaryCategoryID             int                             `json:"primaryCategoryId"`
	Categories                    []Categories                    `json:"categories"`
	ClassID                       int                             `json:"classId"`
	Authors                       []Authors                       `json:"authors"`
	Logo                          Logo                            `json:"logo"`
	Screenshots                   []Screenshots                   `json:"screenshots"`
	MainFileID                    int                             `json:"mainFileId"`
	LatestFiles                   []LatestFiles                   `json:"latestFiles"`
	LatestFilesIndexes            []LatestFilesIndexes            `json:"latestFilesIndexes"`
	LatestEarlyAccessFilesIndexes []LatestEarlyAccessFilesIndexes `json:"latestEarlyAccessFilesIndexes"`
	DateCreated                   time.Time                       `json:"dateCreated"`
	DateModified                  time.Time                       `json:"dateModified"`
	DateReleased                  time.Time                       `json:"dateReleased"`
	AllowModDistribution          bool                            `json:"allowModDistribution"`
	GamePopularityRank            int                             `json:"gamePopularityRank"`
	IsAvailable                   bool                            `json:"isAvailable"`
	ThumbsUpCount                 int                             `json:"thumbsUpCount"`
	Rating                        int                             `json:"rating"`
}
type Pagination struct {
	Index       int `json:"index"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
	TotalCount  int `json:"totalCount"`
}
