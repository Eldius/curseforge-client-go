package client

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/eldius/curseforge-client-go/client/config"
    "github.com/eldius/curseforge-client-go/client/types"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

const (
    modSearchPath = "/v1/mods/search"
    modGetPath    = "/v1/mods/%s"
    //fileGetPath        = "/v1/mods/%s/files/%s"
    fileGetPath        = "/v1/mods/%s/files/%s/download-url"
    gamesListPath      = "/v1/games"
    categoriesListPath = "/v1/categories"
)

// ClientLogger is a logger definition to be used with client
type ClientLogger interface {
    Printf(format string, v ...any)
    Println(v ...any)
}

// Client is the curseforge client itself
type Client struct {
    cfg *config.Config
    log ClientLogger
}

// NewClientWithConfig creates a new Client from a Config instance
func NewClientWithConfig(cfg *config.Config) *Client {
    return &Client{
        cfg: cfg,
        log: log.Default(),
    }
}

// NewClient creates a new client whith default config passing only the API key
func NewClient(apiKey string) *Client {
    return NewClientWithConfig(config.NewDefaultConfig(apiKey))
}

// SetLogger changes the default log implementation (default logger is 'log.Default()')
func (_c *Client) SetLogger(l ClientLogger) {
    _c.log = l
}

// GetGames lists games from API
func (_c *Client) GetGames() (types.GamesResponse, error) {
    var result types.GamesResponse
    c := _c.cfg.NewHTTPClient()

    req, err := _c.cfg.NewGetRequest(c, gamesListPath)
    if err != nil {
        _c.log.Printf("Failed to create request object: %s", err.Error())
        return result, err
    }

    res, err := c.Do(req)
    if err != nil {
        _c.log.Printf("Failed to execute request: %s", err.Error())
        return result, err
    }
    _c.parseResponse(res, &result)

    return result, nil
}

// GetMods lists mods for a game from API
func (_c *Client) GetMods(gameID string) (types.ModsResponse, error) {
    var result types.ModsResponse
    c := _c.cfg.NewHTTPClient()

    req, err := _c.cfg.NewGetRequest(c, modSearchPath)
    if err != nil {
        _c.log.Printf("Failed to create request object: %s", err.Error())
        return result, err
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
        return result, err
    }
    _c.parseResponse(res, &result)

    _c.log.Println("code:", res.StatusCode)

    return result, nil
}

// GetModsByCategory lists mods for a category from API
func (_c *Client) GetModsByCategory(gameID string, modCategorySlug string, searchFilter string) (*types.ModsResponse, error) {
    var result *types.ModsResponse
    c := _c.cfg.NewHTTPClient()

    req, err := _c.cfg.NewGetRequest(c, modSearchPath)
    if err != nil {
        _c.log.Printf("Failed to create request object: %s", err.Error())
        return result, err
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
        return result, err
    }
    if res.StatusCode != http.StatusOK {
        return nil, errors.New("http response error: " + res.Status)
    }

    _c.parseResponse(res, &result)

    return result, nil
}

// GetCategories lists all categories for a game from API
func (_c *Client) GetCategories(gameID string) (*types.ModsResponse, error) {
    var result *types.ModsResponse

    c := _c.cfg.NewHTTPClient()

    req, err := _c.cfg.NewGetRequest(c, categoriesListPath)
    if err != nil {
        _c.log.Printf("Failed to create request object: %s", err.Error())
        return result, err
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
        return result, err
    }
    _c.parseResponse(res, &result)

    if res.StatusCode != http.StatusOK {
        return nil, errors.New("http response error: " + res.Status)
    }

    return result, nil
}

// GetModByID gets mod info by ID from API
func (_c *Client) GetModByID(modID string) (*types.SingleModResult, error) {
    var result *types.SingleModResult
    c := _c.cfg.NewHTTPClient()

    req, err := _c.cfg.NewGetRequest(c, fmt.Sprintf(modGetPath, modID))
    if err != nil {
        _c.log.Printf("Failed to create request object: %s", err.Error())
        return result, err
    }

    if _c.cfg.IsDebug() {
        _c.log.Println(req.URL.String())
    }
    res, err := c.Do(req)
    if err != nil {
        _c.log.Printf("Failed to execute request: %s", err.Error())
        return result, err
    }
    if res.StatusCode != http.StatusOK {
        return nil, errors.New("http response error: " + res.Status)
    }

    _c.parseResponse(res, &result)

    return result, nil
}

// GetFileDownloadURI gets mod info by ID from API
func (_c *Client) GetFileDownloadURI(modID string, fileID string) (*types.GetFileDownloadURLResponse, error) {
    var result *types.GetFileDownloadURLResponse
    c := _c.cfg.NewHTTPClient()

    // /v1/mods/{modId}/files/{fileId}
    req, err := _c.cfg.NewGetRequest(c, fmt.Sprintf(fileGetPath, modID, fileID))
    if err != nil {
        _c.log.Printf("Failed to create request object: %s", err.Error())
        return result, err
    }

    if _c.cfg.IsDebug() {
        _c.log.Println(req.URL.String())
    }
    res, err := c.Do(req)
    if err != nil {
        _c.log.Printf("Failed to execute request: %s", err.Error())
        return result, err
    }
    if res.StatusCode != http.StatusOK {
        return nil, errors.New("http response error: " + res.Status)
    }

    _c.parseResponse(res, &result)

    return result, nil
}

func (_c *Client) debugResponse(r *http.Response) {
    if _c.cfg.IsDebug() {
        reader := r.Body
        defer func() {
            err := reader.Close()
            if err != nil {
                _c.log.Println("Failed to close reader:", err.Error())
            }
        }()
        b, err := ioutil.ReadAll(reader)
        if err != nil {
            _c.log.Printf("Failed to execute request: %s", err.Error())
            return
        }
        msg := "---\nheaders:\n"
        for k, v := range r.Header {
            msg += fmt.Sprintf(" - %s: [%s]\n", k, strings.Join(v, ", "))
        }
        msg += fmt.Sprintf("response:\n%s\n---", string(b))
        _c.log.Println(msg)
        _c.log.Println()
        r.Body = io.NopCloser(bytes.NewReader(b))
    }
}

func (_c *Client) parseResponse(r *http.Response, result interface{}) {
    _c.debugResponse(r)
    defer func() {
        _ = r.Body.Close()
    }()
    _ = json.NewDecoder(r.Body).Decode(result)
}
