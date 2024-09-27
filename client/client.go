package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eldius/curseforge-client-go/client/config"
	"github.com/eldius/curseforge-client-go/client/types"
	"io"
	"net/http"
)

const (
	modSearchPath      = "/v1/mods/search"
	modGetPath         = "/v1/mods/%s"
	fileGetPath        = "/v1/mods/%s/files/%s/download-url"
	gamesListPath      = "/v1/games"
	categoriesListPath = "/v1/categories"
)

// Logger is a logger definition to be used with client
type Logger interface {
	Printf(format string, v ...any)
	Println(v ...any)
	DebugRequest(res *http.Response)
}

type noopLogger struct {
	Logger
}

func (l noopLogger) Printf(_ string, _ ...any) {
}

func (l noopLogger) Println(_ ...any) {
}

func (l noopLogger) DebugRequest(_ *http.Response) {
}

// Client is the curseforge client itself
type Client struct {
	cfg *config.Config
	log Logger
}

// NewClientWithConfig creates a new Client from a Config instance
func NewClientWithConfig(cfg *config.Config) *Client {
	return &Client{
		cfg: cfg,
		log: noopLogger{},
	}
}

// NewClient creates a new client whith default config passing only the API key
func NewClient(apiKey string) *Client {
	return NewClientWithConfig(config.NewConfig(apiKey))
}

// SetLogger changes the default log implementation (default logger is 'log.Default()')
func (_c *Client) SetLogger(l Logger) {
	_c.log = l
}

// GetGames lists games from API
func (_c *Client) GetGames() (types.GamesResponse, error) {
	var result types.GamesResponse
	c := _c.cfg.NewHTTPClient()

	req, err := _c.cfg.NewGetRequest(c, gamesListPath)
	if err != nil {
		_c.log.Printf("Failed to create request object: %s", err.Error())
		return result, types.Wrap(err, "failed to create request object", -1)
	}

	res, err := c.Do(req)
	if err != nil {
		_c.log.Printf("Failed to execute request: %s", err.Error())
		return result, types.Wrap(err, "failed to execute request", -1)
	}
	if err := parseResponse(res, &result); err != nil {
		err = fmt.Errorf("parsing API response")
		return result, err
	}

	return result, nil
}

// GetMods lists mods for a game from API
func (_c *Client) GetMods(f ModFilter) (*types.ModsResponse, error) {
	var result types.ModsResponse
	c := _c.cfg.NewHTTPClient()

	req, err := _c.cfg.NewGetRequest(c, modSearchPath)
	if err != nil {
		_c.log.Printf("Failed to create request object: %s", err.Error())
		return &result, types.Wrap(err, "failed to create request object", -1)
	}

	req.URL.RawQuery = f.QueryString()
	res, err := c.Do(req)
	if err != nil {
		_c.log.Printf("Failed to execute request: %s", err.Error())
		return &result, types.Wrap(err, "failed to execute request", -1)
	}

	if err := parseResponse(res, &result); err != nil {
		err = fmt.Errorf("parsing API response")
		return &result, err
	}

	_c.log.Println("code:", res.StatusCode)

	return &result, nil
}

// GetModsByCategory lists mods for a category from API
func (_c *Client) GetModsByCategory(gameID string, modCategorySlug string, searchFilter string) (*types.ModsResponse, error) {
	var result types.ModsResponse
	c := _c.cfg.NewHTTPClient()

	req, err := _c.cfg.NewGetRequest(c, modSearchPath)
	if err != nil {
		_c.log.Printf("Failed to create request object: %s", err.Error())
		return nil, types.Wrap(err, "failed to create request object", -1)
	}

	q := req.URL.Query()
	q.Add("gameId", gameID)
	q.Add("categoryId", modCategorySlug)
	q.Add("searchFilter", searchFilter)
	req.URL.RawQuery = q.Encode()
	if _c.cfg.IsDebug() {
		_c.log.Println(req.URL.String())
	}
	res, err := c.Do(req)
	if err != nil {
		_c.log.Printf("Failed to execute request: %s", err.Error())
		return nil, types.Wrap(err, "failed to execute request", -1)
	}
	if res.StatusCode != http.StatusOK {
		return nil, types.Wrap(errors.New("http response error"), "api request error", res.StatusCode)
	}

	if err := parseResponse(res, &result); err != nil {
		err = fmt.Errorf("parsing API response")
		return nil, err
	}

	return &result, nil
}

// GetCategories lists all categories for a game from API
func (_c *Client) GetCategories(gameID string) (*types.ModsResponse, error) {
	var result types.ModsResponse

	c := _c.cfg.NewHTTPClient()

	req, err := _c.cfg.NewGetRequest(c, categoriesListPath)
	if err != nil {
		_c.log.Printf("Failed to create request object: %s", err.Error())
		return nil, err
	}

	q := req.URL.Query()
	q.Add("gameId", gameID)
	req.URL.RawQuery = q.Encode()
	if _c.cfg.IsDebug() {
		_c.log.Println(req.URL.String())
	}
	res, err := c.Do(req)
	if err != nil {
		_c.log.Printf("Failed to execute request: %s", err.Error())
		return nil, types.Wrap(err, types.ErrRequestErrorMsg, 0)
	}
	if err := parseResponse(res, &result); err != nil {
		err = fmt.Errorf("parsing API response")
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http response error: " + res.Status)
	}

	return &result, nil
}

// GetModByID gets mod info by ID from API
func (_c *Client) GetModByID(modID string) (*types.SingleModResult, error) {
	var result types.SingleModResult
	c := _c.cfg.NewHTTPClient()

	req, err := _c.cfg.NewGetRequest(c, fmt.Sprintf(modGetPath, modID))
	if err != nil {
		_c.log.Printf("Failed to create request object: %s", err.Error())
		return nil, err
	}

	if _c.cfg.IsDebug() {
		_c.log.Println(req.URL.String())
	}
	res, err := c.Do(req)
	if err != nil {
		_c.log.Printf("Failed to execute request: %s", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http response error: " + res.Status)
	}

	if err := parseResponse(res, &result); err != nil {
		err = fmt.Errorf("parsing API response")
		return nil, err
	}

	return &result, nil
}

// GetFileDownloadURI gets mod info by ID from API
func (_c *Client) GetFileDownloadURI(modID string, fileID string) (*types.GetFileDownloadURLResponse, error) {
	var result types.GetFileDownloadURLResponse
	c := _c.cfg.NewHTTPClient()

	// /v1/mods/{modId}/files/{fileId}
	req, err := _c.cfg.NewGetRequest(c, fmt.Sprintf(fileGetPath, modID, fileID))
	if err != nil {
		_c.log.Printf("Failed to create request object: %s", err.Error())
		return nil, err
	}

	if _c.cfg.IsDebug() {
		_c.log.Println(req.URL.String())
	}
	res, err := c.Do(req)
	if err != nil {
		_c.log.Printf("Failed to execute request: %s", err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http response error: " + res.Status)
	}

	if err := parseResponse(res, &result); err != nil {
		err = fmt.Errorf("parsing API response")
		return nil, err
	}

	return &result, nil
}

func parseResponse(res *http.Response, result types.CurseforgeAPIResponse) error {
	defer func() {
		_ = res.Body.Close()
	}()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if err := json.Unmarshal(b, result); err != nil {
		return fmt.Errorf("parsing api response: %w", err)
	}

	result.SetResponse(string(b))

	return nil
}
