package types

import (
	"slices"
	"sort"
	"time"
)

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

const (
	RelationTypeEmbeddedLibrary    RelationType = 1
	RelationTypeOptionalDependency RelationType = 2
	RelationTypeRequiredDependency RelationType = 3
	RelationTypeTool               RelationType = 4
	RelationTypeIncompatible       RelationType = 5
	RelationTypeInclude            RelationType = 6
)

const (
	HashAlgoSHA1 HashAlgo = 1
	HashAlgoMD5  HashAlgo = 2
)

const (
	FileReleaseTypeRelease FileReleaseType = 1
	FileReleaseTypeBeta    FileReleaseType = 2
	FileReleaseTypeAlpha   FileReleaseType = 3
)

const (
	FileStatusProcessing         = 1
	FileStatusChangesRequired    = 2
	FileStatusUnderReview        = 3
	FileStatusApproved           = 4
	FileStatusRejected           = 5
	FileStatusMalwareDetected    = 6
	FileStatusDeleted            = 7
	FileStatusArchived           = 8
	FileStatusTesting            = 9
	FileStatusReleased           = 10
	FileStatusReadyForReview     = 11
	FileStatusDeprecated         = 12
	FileStatusBaking             = 13
	FileStatusAwaitingPublishing = 14
	FileStatusFailedPublishing   = 15
)

type FileStatus int

func (f FileStatus) String() string {
	switch f {
	case FileStatusProcessing:
		return "processing"
	case FileStatusChangesRequired:
		return "changes-required"
	case FileStatusUnderReview:
		return "under-review"
	case FileStatusApproved:
		return "approved"
	case FileStatusRejected:
		return "rejected"
	case FileStatusMalwareDetected:
		return "malware-detected"
	case FileStatusDeleted:
		return "deleted"
	case FileStatusArchived:
		return "archived"
	case FileStatusTesting:
		return "testing"
	case FileStatusReleased:
		return "released"
	case FileStatusReadyForReview:
		return "ready-for-review"
	case FileStatusDeprecated:
		return "deprecated"
	case FileStatusBaking:
		return "baking"
	case FileStatusAwaitingPublishing:
		return "awaiting-publishing"
	case FileStatusFailedPublishing:
		return "failed-publishing"
	default:
		return "unknown"
	}
}

type FileReleaseType int

func (f FileReleaseType) String() string {
	switch f {
	case FileReleaseTypeRelease:
		return "Release"
	case FileReleaseTypeBeta:
		return "Beta"
	case FileReleaseTypeAlpha:
		return "Alpha"
	default:
		return "Unknown"
	}
}

type HashAlgo int

func (h HashAlgo) String() string {
	switch h {
	case HashAlgoSHA1:
		return "SHA1"
	case HashAlgoMD5:
		return "MD5"
	default:
		return "Unknown"
	}
}

type RelationType int

func (r RelationType) String() string {
	switch r {
	case RelationTypeEmbeddedLibrary:
		return "embedded library"
	case RelationTypeOptionalDependency:
		return "optional"
	case RelationTypeRequiredDependency:
		return "required"
	case RelationTypeTool:
		return "tool"
	case RelationTypeIncompatible:
		return "incompatible"
	case RelationTypeInclude:
		return "include"
	default:
		return "unknown"
	}
}

type ModLoaderType int

func (m ModLoaderType) String() string {
	switch m {
	case ModLoaderTypeAny:
		return "any"
	case ModLoaderTypeForge:
		return "forge"
	case ModLoaderTypeCauldron:
		return "cauldron"
	case ModLoaderTypeLiteLoader:
		return "liteloader"
	case ModLoaderTypeFabric:
		return "fabric"
	case ModLoaderTypeQuilt:
		return "quilt"
	case ModLoaderTypeNeoForge:
		return "neoforge"
	default:
		return "unknown"
	}
}

func (t ModLoaderType) Value() int {
	return int(t)
}

type ModStatus int

func (m ModStatus) String() string {
	switch m {
	case ModStatusNew:
		return "new"
	case ModStatusChangesRequired:
		return "changes-required"
	case ModStatusUnderReview:
		return "under-review"
	case ModStatusApproved:
		return "approved"
	case ModStatusRejected:
		return "rejected"
	case ModStatusChangesMade:
		return "changes-made"
	case ModStatusInactive:
		return "inactive"
	case ModStatusAbandoned:
		return "abandoned"
	case ModStatusDeleted:
		return "deleted"
	case ModStatusUnderSoftReview:
		return "under-soft-review"
	default:
		return "unknown"
	}
}

type GameVersionsResponse struct {
	RawResponse
	Data GameVersions `json:"data"`
}

type GameVersions []GameVersion

func (v GameVersions) Sort() GameVersions {
	sort.Slice(v, func(i, j int) bool {
		return v[i].Name < v[j].Name
	})
	return v
}

type GameVersion struct {
	ID   int64  `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type MinecraftVersionsResponse struct {
	RawResponse
	Data MinecraftVersions `json:"data"`
}

type MinecraftVersions []MinecraftVersion

func (v MinecraftVersions) Sort() MinecraftVersions {
	sort.Slice(v, func(i, j int) bool {
		return v[i].DateModified.After(v[j].DateModified)
	})
	return v
}

type MinecraftVersion struct {
	ID                    int64     `json:"id"`
	GameVersionID         int64     `json:"gameVersionId"`
	VersionString         string    `json:"versionString"`
	JarDownloadUrl        string    `json:"jarDownloadUrl"`
	JsonDownloadUrl       string    `json:"jsonDownloadUrl"`
	Approved              bool      `json:"approved"`
	DateModified          time.Time `json:"dateModified"`
	GameVersionTypeID     int64     `json:"gameVersionTypeId"`
	GameVersionStatus     int64     `json:"gameVersionStatus"`
	GameVersionTypeStatus int64     `json:"gameVersionTypeStatus"`
}

type MinecraftModLoadersResponse struct {
	RawResponse
	Data MinecraftModLoaders `json:"data"`
}

type MinecraftModLoaders []MinecraftModLoader

func (v MinecraftModLoaders) Sort() MinecraftModLoaders {
	sort.Slice(v, func(i, j int) bool {
		return v[i].DateModified.After(v[j].DateModified)
	})
	return v
}

type MinecraftModLoader struct {
	Name         string    `json:"name"`
	GameVersion  string    `json:"gameVersion"`
	Latest       bool      `json:"latest"`
	Recommended  bool      `json:"recommended"`
	DateModified time.Time `json:"dateModified"`
	Type         int       `json:"type"`
}

type ModsResponse struct {
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
	ID               int64     `json:"id"`
	GameID           int64     `json:"gameId"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	URL              string    `json:"url"`
	IconURL          string    `json:"iconUrl"`
	DateModified     time.Time `json:"dateModified"`
	IsClass          bool      `json:"isClass"`
	ClassID          int64     `json:"classId"`
	ParentCategoryID int64     `json:"parentCategoryId"`
	DisplayIndex     int64     `json:"displayIndex"`
}

type Authors struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Logo struct {
	ID           int64  `json:"id"`
	ModID        int64  `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

type Screenshots struct {
	ID           int64  `json:"id"`
	ModID        int64  `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

type Hashes struct {
	Value string   `json:"value"`
	Algo  HashAlgo `json:"algo"`
}

type SortableGameVersions struct {
	GameVersionName        string    `json:"gameVersionName"`
	GameVersionPadded      string    `json:"gameVersionPadded"`
	GameVersion            string    `json:"gameVersion"`
	GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
	GameVersionTypeID      int64     `json:"gameVersionTypeId"`
}

type Dependencies struct {
	ModID        int64        `json:"modId"`
	RelationType RelationType `json:"relationType"`
}

type Modules struct {
	Name        string `json:"name"`
	Fingerprint int64  `json:"fingerprint"`
}

type Files []File

type File struct {
	ID                   int64                  `json:"id"`
	GameID               int64                  `json:"gameId"`
	ModID                int64                  `json:"modId"`
	IsAvailable          bool                   `json:"isAvailable"`
	DisplayName          string                 `json:"displayName"`
	FileName             string                 `json:"fileName"`
	ReleaseType          FileReleaseType        `json:"releaseType"`
	FileStatus           FileStatus             `json:"fileStatus"`
	Hashes               []Hashes               `json:"hashes"`
	FileDate             time.Time              `json:"fileDate"`
	FileLength           int64                  `json:"fileLength"`
	DownloadCount        int64                  `json:"downloadCount"`
	FileSizeOnDisk       int64                  `json:"fileSizeOnDisk"`
	DownloadURL          string                 `json:"downloadUrl"`
	GameVersions         []string               `json:"gameVersions"`
	SortableGameVersions []SortableGameVersions `json:"sortableGameVersions"`
	Dependencies         []Dependencies         `json:"dependencies"`
	ExposeAsAlternative  bool                   `json:"exposeAsAlternative"`
	ParentProjectFileID  int64                  `json:"parentProjectFileId"`
	AlternateFileID      int64                  `json:"alternateFileId"`
	IsServerPack         bool                   `json:"isServerPack"`
	ServerPackFileID     int64                  `json:"serverPackFileId"`
	IsEarlyAccessContent bool                   `json:"isEarlyAccessContent"`
	EarlyAccessEndDate   time.Time              `json:"earlyAccessEndDate"`
	FileFingerprint      int64                  `json:"fileFingerprint"`
	Modules              []Modules              `json:"modules"`
}

type LatestFilesIndexes struct {
	GameVersion       string        `json:"gameVersion"`
	FileID            int64         `json:"fileId"`
	Filename          string        `json:"filename"`
	ReleaseType       int           `json:"releaseType"`
	GameVersionTypeID int           `json:"gameVersionTypeId"`
	ModLoader         ModLoaderType `json:"modLoader"`
}

type LatestEarlyAccessFilesIndexes struct {
	GameVersion       string `json:"gameVersion"`
	FileID            int64  `json:"fileId"`
	Filename          string `json:"filename"`
	ReleaseType       int    `json:"releaseType"`
	GameVersionTypeID int    `json:"gameVersionTypeId"`
	ModLoader         int    `json:"modLoader"`
}

type ModData struct {
	ID                            int64                           `json:"id"`
	GameID                        int                             `json:"gameId"`
	Name                          string                          `json:"name"`
	Slug                          string                          `json:"slug"`
	Links                         Links                           `json:"links"`
	Summary                       string                          `json:"summary"`
	Status                        ModStatus                       `json:"status"`
	DownloadCount                 int                             `json:"downloadCount"`
	IsFeatured                    bool                            `json:"isFeatured"`
	PrimaryCategoryID             int                             `json:"primaryCategoryId"`
	Categories                    []Categories                    `json:"categories"`
	ClassID                       int                             `json:"classId"`
	Authors                       []Authors                       `json:"authors"`
	Logo                          Logo                            `json:"logo"`
	Screenshots                   []Screenshots                   `json:"screenshots"`
	MainFileID                    int                             `json:"mainFileId"`
	LatestFiles                   Files                           `json:"latestFiles"`
	LatestFilesIndexes            []LatestFilesIndexes            `json:"latestFilesIndexes"`
	LatestEarlyAccessFilesIndexes []LatestEarlyAccessFilesIndexes `json:"latestEarlyAccessFilesIndexes"`
	DateCreated                   time.Time                       `json:"dateCreated"`
	DateModified                  time.Time                       `json:"dateModified"`
	DateReleased                  time.Time                       `json:"dateReleased"`
	AllowModDistribution          bool                            `json:"allowModDistribution"`
	GamePopularityRank            int64                           `json:"gamePopularityRank"`
	IsAvailable                   bool                            `json:"isAvailable"`
	ThumbsUpCount                 int64                           `json:"thumbsUpCount"`
	Rating                        int64                           `json:"rating"`
}

type Pagination struct {
	Index       int64 `json:"index"`
	PageSize    int64 `json:"pageSize"`
	ResultCount int64 `json:"resultCount"`
	TotalCount  int64 `json:"totalCount"`
}

// GetLatestFile returns the latest mod file
func (m *ModData) GetLatestFile() *File {
	if len(m.LatestFiles) < 1 {
		return nil
	}
	return &m.LatestFiles[0]
}

// GetLatestFileGameVersions returns the latest file version data
func (m *ModData) GetLatestFileGameVersions() []string {
	f := m.GetLatestFile()
	if f == nil {
		return nil
	}

	return f.GameVersions
}

// GetLatestFileByGameVersion returns the latest file version data
func (m *ModData) GetLatestFileByGameVersion(gv string) *File {

	for _, f := range m.LatestFiles {
		if slices.Contains(f.GameVersions, gv) {
			return &f
		}
	}

	return nil
}
