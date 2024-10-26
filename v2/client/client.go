package client

import (
	"encoding/json"
	"fmt"
	"github.com/eldius/curseforge-client-go/v2/client/config"
	"github.com/eldius/curseforge-client-go/v2/client/types"
	"io"
	"net/http"
)

const (
	getGameVersionsPath    = "/v2/games/%s/versions?sortDescending=true"
	getMinecraftVersions   = "/v1/minecraft/version"
	getMinecraftModLoaders = "/v1/minecraft/modloader"
)

type CurseClient interface {
	GetGameVersions(gameID string) (versions *types.GameVersionsResponse, err error)
	GetMinecraftVersions(...MinecraftVersionsQueryOption) (versions *types.MinecraftVersionsResponse, err error)
	GetMinecraftModLoaders(...MinecraftModLoadersQueryOption) (versions *types.MinecraftModLoadersResponse, err error)
}

type CurseOptions struct {
	apiKey   string
	endpoint string
}

type CurseOption func(*CurseOptions)

type curseClient struct {
	CurseClient
	opt config.Config
	c   *http.Client
}

// WithCurseApiKey sets up Curseforge API key
func WithCurseApiKey(apiKey string) CurseOption {
	return func(o *CurseOptions) {
		o.apiKey = apiKey
	}
}

func WithCustomEndpoint(endpoint string) CurseOption {
	return func(o *CurseOptions) {
		o.endpoint = endpoint
	}
}

// NewCurseClient creates a new Curseforge client
func NewCurseClient(apiKey string, opts ...config.CfgFunc) CurseClient {
	opt := config.NewConfig(apiKey, opts...)
	return &curseClient{
		opt: *opt,
		c:   &http.Client{},
	}
}

func (c *curseClient) GetGameVersions(gameID string) (versions *types.GameVersionsResponse, err error) {
	req, err := c.opt.NewGetRequest(c.buildRequestPath(fmt.Sprintf(getGameVersionsPath, gameID), ApiQueryParams{}))
	if err != nil {
		return nil, fmt.Errorf("creating get game versions request: %w", err)
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing get game versions request: %w", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode/100 != 2 {
		b, _ := io.ReadAll(res.Body)
		return nil, types.Wrap(
			fmt.Errorf("get game versions request failed with status code %d", res.StatusCode),
			string(b),
			res.StatusCode,
		)
	}
	var gv types.GameVersionsResponse
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading get game versions response: %w", err)
	}
	if err := json.Unmarshal(b, &gv); err != nil {
		return nil, fmt.Errorf("decoding get game versions response: %w", err)
	}
	gv.RawBody = string(b)
	return &gv, nil
}

func (c *curseClient) GetMinecraftVersions(opts ...MinecraftVersionsQueryOption) (versions *types.MinecraftVersionsResponse, err error) {
	q := ApiQueryParams{}
	for _, o := range opts {
		o(q)
	}
	req, err := c.opt.NewGetRequest(c.buildRequestPath(getMinecraftVersions, q))
	if err != nil {
		return nil, fmt.Errorf("creating get game versions request: %w", err)
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing get minecraft versions request: %w", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode/100 != 2 {
		b, _ := io.ReadAll(res.Body)
		return nil, types.Wrap(
			fmt.Errorf("get game versions request failed with status code %d", res.StatusCode),
			string(b),
			res.StatusCode,
		)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading get minecraft versions response: %w", err)
	}
	var mv types.MinecraftVersionsResponse
	if err := json.Unmarshal(b, &mv); err != nil {
		return nil, fmt.Errorf("decoding get minecraft versions response: %w", err)
	}
	mv.RawBody = string(b)

	return &mv, nil
}

func (c *curseClient) GetMinecraftModLoaders(opts ...MinecraftModLoadersQueryOption) (versions *types.MinecraftModLoadersResponse, err error) {
	q := ApiQueryParams{}
	for _, o := range opts {
		o(q)
	}

	req, err := c.opt.NewGetRequest(c.buildRequestPath(getMinecraftModLoaders, q))
	if err != nil {
		return nil, fmt.Errorf("creating get game versions request: %w", err)
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing get minecraft mod loaders request: %w", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()
	if res.StatusCode/100 != 2 {
		b, _ := io.ReadAll(res.Body)
		return nil, types.Wrap(
			fmt.Errorf("get game versions request failed with status code %d", res.StatusCode),
			string(b),
			res.StatusCode,
		)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading get minecraft mod loaders response: %w", err)
	}
	var mv types.MinecraftModLoadersResponse
	if err := json.Unmarshal(b, &mv); err != nil {
		return nil, fmt.Errorf("decoding get minecraft mod loaders response: %w", err)
	}
	mv.RawBody = string(b)

	return &mv, nil
}

func (c *curseClient) buildRequestPath(path string, q ApiQueryParams) string {
	return fmt.Sprintf("%s?%s", path, q.QueryString())
}
