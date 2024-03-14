package types

import (
	"strings"
	"time"
)

// GamesResponse is the games search response
type GamesResponse struct {
	Data       []Game     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// Assets is an asset representation
type Assets struct {
	IconURL  string `json:"iconUrl"`
	TileURL  string `json:"tileUrl"`
	CoverURL string `json:"coverUrl"`
}

// Game is the game info
type Game struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	DateModified time.Time `json:"dateModified"`
	Assets       Assets    `json:"assets"`
	Status       int       `json:"status"`
	APIStatus    int       `json:"apiStatus"`
}

// Pagination is the API response pagination info
type Pagination struct {
	Index       int `json:"index"`
	PageSize    int `json:"pageSize"`
	ResultCount int `json:"resultCount"`
	TotalCount  int `json:"totalCount"`
}

// ModsResponse is the mod search response representation
type ModsResponse struct {
	Data       []ModData  `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// Links is the mod links representation
type Links struct {
	WebsiteURL string `json:"websiteUrl"`
	WikiURL    string `json:"wikiUrl"`
	IssuesURL  string `json:"issuesUrl"`
	SourceURL  string `json:"sourceUrl"`
}

// Categories represents mod categories
type Categories struct {
	ID               int64     `json:"id"`
	GameID           int64     `json:"gameId"`
	Name             string    `json:"name"`
	Slug             string    `json:"slug"`
	URL              string    `json:"url"`
	IconURL          string    `json:"iconUrl"`
	DateModified     time.Time `json:"dateModified"`
	IsClass          bool      `json:"isClass"`
	ClassID          int       `json:"classId"`
	ParentCategoryID int       `json:"parentCategoryId"`
}

// Authors is the Author info
type Authors struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Logo is the mod logo representation
type Logo struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

// Screenshots represents a screenshot data
type Screenshots struct {
	ID           int    `json:"id"`
	ModID        int    `json:"modId"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
}

// Hashes is a file hash info
type Hashes struct {
	Value string `json:"value"`
	Algo  int    `json:"algo"`
}

// SortableGameVersions have game version infos
type SortableGameVersions struct {
	GameVersionName        string    `json:"gameVersionName"`
	GameVersionPadded      string    `json:"gameVersionPadded"`
	GameVersion            string    `json:"gameVersion"`
	GameVersionReleaseDate time.Time `json:"gameVersionReleaseDate"`
	GameVersionTypeID      int       `json:"gameVersionTypeId"`
}

// Modules represents a module(?)
type Modules struct {
	Name        string `json:"name"`
	Fingerprint int64  `json:"fingerprint"`
}

// File represents a file
type File struct {
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
	ServerPackFileID     int                    `json:"serverPackFileId"`
	FileFingerprint      int64                  `json:"fileFingerprint"`
	Modules              []Modules              `json:"modules"`
}

// FileIndexes represents file indexes
type FileIndexes struct {
	GameVersion       string `json:"gameVersion"`
	FileID            int    `json:"fileId"`
	Filename          string `json:"filename"`
	ReleaseType       int    `json:"releaseType"`
	GameVersionTypeID int    `json:"gameVersionTypeId"`
	ModLoader         int    `json:"modLoader"`
}

// ModData is a mod info representation
type ModData struct {
	ID                   int64         `json:"id"`
	GameID               int64         `json:"gameId"`
	Name                 string        `json:"name"`
	Slug                 string        `json:"slug"`
	Links                Links         `json:"links"`
	Summary              string        `json:"summary"`
	Status               int           `json:"status"`
	DownloadCount        float64       `json:"downloadCount"`
	IsFeatured           bool          `json:"isFeatured"`
	PrimaryCategoryID    int           `json:"primaryCategoryId"`
	Categories           []Categories  `json:"categories"`
	ClassID              int64         `json:"classId"`
	Authors              []Authors     `json:"authors"`
	Logo                 Logo          `json:"logo"`
	Screenshots          []Screenshots `json:"screenshots"`
	MainFileID           int           `json:"mainFileId"`
	LatestFiles          []File        `json:"latestFiles"`
	LatestFilesIndexes   []FileIndexes `json:"latestFilesIndexes"`
	DateCreated          time.Time     `json:"dateCreated"`
	DateModified         time.Time     `json:"dateModified"`
	DateReleased         time.Time     `json:"dateReleased"`
	AllowModDistribution interface{}   `json:"allowModDistribution"`
	GamePopularityRank   int           `json:"gamePopularityRank"`
}

// GetFileDownloadURLResponse file download URL
type GetFileDownloadURLResponse struct {
	URL string `json:"data"`
}

// TranslateStatus translates file status to text
func (f *File) TranslateStatus() string {
	return translateStatus(f.FileStatus, FileStatusMap)
}

// TranslateStatus translates file status to text
func (g *Game) TranslateStatus() string {
	//return translateStatus(g.Status)
	return "Not implemented yet."
}

// TranslateStatus translates file status to text
func (m *ModData) TranslateStatus() string {
	return translateStatus(m.Status, ModStatusMap)
}

func translateStatus(s int, m map[int]string) string {
	res, ok := m[s]
	if !ok {
		return "UNKNOWN"
	}
	return res
}

// GetLatestFile returns the latest mod file
func (m *ModData) GetLatestFile() *File {
	if len(m.LatestFiles) < 1 {
		return nil
	}
	return &m.LatestFiles[0]
}

// GetLatestFileGameVersions returns the latest file version data
func (m *ModData) GetLatestFileGameVersions() string {
	f := m.GetLatestFile()
	if f == nil {
		return ""
	}

	return strings.Join(f.GameVersions, ", ")
}

// SingleModResult is the result of APIs single mod data response
type SingleModResult struct {
	Data ModData `json:"data"`
}
