package curseforge

import "time"

type GamesResponse struct {
	Data       []Game     `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Assets struct {
	IconURL  string `json:"iconUrl"`
	TileURL  string `json:"tileUrl"`
	CoverURL string `json:"coverUrl"`
}
type Game struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	DateModified time.Time `json:"dateModified"`
	Assets       Assets    `json:"assets"`
	Status       int       `json:"status"`
	APIStatus    int       `json:"apiStatus"`
}
type Pagination struct {
	Index       int `json:"index"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
	TotalCount  int `json:"totalCount"`
}

type ModsResponse struct {
	Data       []ModData  `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Links struct {
	WebsiteURL string      `json:"websiteUrl"`
	WikiURL    string      `json:"wikiUrl"`
	IssuesURL  interface{} `json:"issuesUrl"`
	SourceURL  interface{} `json:"sourceUrl"`
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
type Modules struct {
	Name        string `json:"name"`
	Fingerprint int64  `json:"fingerprint"`
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
	DownloadURL          string                 `json:"downloadUrl"`
	GameVersions         []string               `json:"gameVersions"`
	SortableGameVersions []SortableGameVersions `json:"sortableGameVersions"`
	Dependencies         []interface{}          `json:"dependencies"`
	AlternateFileID      int                    `json:"alternateFileId"`
	IsServerPack         bool                   `json:"isServerPack"`
	FileFingerprint      int64                  `json:"fileFingerprint"`
	Modules              []Modules              `json:"modules"`
}
type LatestFilesIndexes struct {
	GameVersion       string `json:"gameVersion"`
	FileID            int    `json:"fileId"`
	Filename          string `json:"filename"`
	ReleaseType       int    `json:"releaseType"`
	GameVersionTypeID int    `json:"gameVersionTypeId"`
	ModLoader         int    `json:"modLoader"`
}
type ModData struct {
	ID                   int                  `json:"id"`
	GameID               int                  `json:"gameId"`
	Name                 string               `json:"name"`
	Slug                 string               `json:"slug"`
	Links                Links                `json:"links"`
	Summary              string               `json:"summary"`
	Status               int                  `json:"status"`
	DownloadCount        float64              `json:"downloadCount"`
	IsFeatured           bool                 `json:"isFeatured"`
	PrimaryCategoryID    int                  `json:"primaryCategoryId"`
	Categories           []Categories         `json:"categories"`
	ClassID              int                  `json:"classId"`
	Authors              []Authors            `json:"authors"`
	Logo                 Logo                 `json:"logo"`
	Screenshots          []Screenshots        `json:"screenshots"`
	MainFileID           int                  `json:"mainFileId"`
	LatestFiles          []LatestFiles        `json:"latestFiles"`
	LatestFilesIndexes   []LatestFilesIndexes `json:"latestFilesIndexes"`
	DateCreated          time.Time            `json:"dateCreated"`
	DateModified         time.Time            `json:"dateModified"`
	DateReleased         time.Time            `json:"dateReleased"`
	AllowModDistribution interface{}          `json:"allowModDistribution"`
	GamePopularityRank   int                  `json:"gamePopularityRank"`
}

func (f *LatestFiles) TranslateStatus() string {
	return translateStatus(f.FileStatus)
}

func (g *Game) TranslateStatus() string {
	return translateStatus(g.Status)
}

func (m *ModData) TranslateStatus() string {
	s, ok := ModStatusMap[m.Status]
	if !ok {
		return "UNKNOWN"
	}
	return s
}

func translateStatus(s int) string {
	switch s {
	case CoreStatusDraft:
		return CoreStatusTextDraft
	case CoreStatusTest:
		return CoreStatusTextTest
	case CoreStatusPendingReview:
		return CoreStatusTextPendingReview
	case CoreStatusRejected:
		return CoreStatusTextRejected
	case CoreStatusApproved:
		return CoreStatusTextApproved
	case CoreStatusLive:
		return CoreStatusTextLive
	default:
		return "UNKNOWN"
	}
}
