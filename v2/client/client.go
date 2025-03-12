package client

import (
	"encoding/json"
	"fmt"
	"github.com/eldius/curseforge-client-go/v2/client/types"
	"io"
	"log"
	"net/http"
)

const (
	getGameVersionsPath    = "/v2/games/%s/versions?sortDescending=true"
	getMinecraftVersions   = "/v1/minecraft/version"
	getMinecraftModLoaders = "/v1/minecraft/modloader"
	minecraftModSearchPath = "/v1/mods/search"
)

type CurseClient interface {
	GetGameVersions(gameID string) (versions *types.GameVersionsResponse, err error)
	GetMinecraftVersions(...MinecraftVersionsQueryOption) (versions *types.MinecraftVersionsResponse, err error)
	GetMinecraftModLoaders(...MinecraftModLoadersQueryOption) (versions *types.MinecraftModLoadersResponse, err error)
	GetMods(opts ...ModsQueryOption) (*types.ModsResponse, error)
	GetModsByIDs(req *GetModsByIdsListRequest, opts ...ModsQueryOption) (*types.ModsResponse, error)
}

type curseClient struct {
	CurseClient
	opt Config
	c   *http.Client
}

// NewCurseClient creates a new Curseforge client
func NewCurseClient(apiKey string, opts ...CfgFunc) CurseClient {
	opt := NewConfig(apiKey, opts...)
	return &curseClient{
		opt: *opt,
		c:   opt.NewHTTPClient(),
	}
}

func (c *curseClient) GetGameVersions(gameID string) (versions *types.GameVersionsResponse, err error) {
	log := c.opt.log
	url := c.buildRequestPath(fmt.Sprintf(getGameVersionsPath, gameID), ApiQueryParams{})
	if c.opt.debug {
		log.Printf("GetGameVersions.Begin id: %s (url: %s)", gameID, url)
	}
	req, err := c.opt.NewGetRequest(url)
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
	var gv types.GameVersionsResponse
	if err := parseResponse(res, "get game versions", c.opt.debug, &gv); err != nil {
		return nil, err
	}
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
	var mv types.MinecraftVersionsResponse
	if err := parseResponse(res, "get minecraft versions", c.opt.debug, &mv); err != nil {
		return nil, err
	}

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
	var mv types.MinecraftModLoadersResponse
	if err := parseResponse(res, "get minecraft mod loaders", c.opt.debug, &mv); err != nil {
		return nil, err
	}

	return &mv, nil
}

// GetMods lists mods for a game from API
func (c *curseClient) GetMods(opts ...ModsQueryOption) (*types.ModsResponse, error) {
	q := ApiQueryParams{}
	for _, o := range opts {
		o(q)
	}
	var result types.ModsResponse
	req, err := c.opt.NewGetRequest(c.buildRequestPath(minecraftModSearchPath, q))
	if err != nil {
		err = fmt.Errorf("creating get minecraft mods request: %w", err)
		return &result, types.Wrap(err, "failed to create request object", -1)
	}

	res, err := c.c.Do(req)
	if err != nil {
		err = fmt.Errorf("executing get minecraft mods request: %w", err)
		return &result, types.Wrap(err, "failed to execute request", -1)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var mv types.ModsResponse
	if err := parseResponse(res, "get game versions", c.opt.debug, &mv); err != nil {
		return nil, fmt.Errorf("parsing get minecraft mods response: %w", err)
	}

	return &mv, nil
}

func (c *curseClient) GetModsByIDs(filter *GetModsByIdsListRequest, opts ...ModsQueryOption) (*types.ModsResponse, error) {
	q := ApiQueryParams{}
	for _, o := range opts {
		o(q)
	}
	var result types.ModsResponse
	req, err := c.opt.NewPostRequest(fmt.Sprintf("%s/v1/mods", minecraftModSearchPath), filter)
	if err != nil {
		err = fmt.Errorf("creating get minecraft mods request: %w", err)
		return &result, types.Wrap(err, "failed to create request object", -1)
	}

	res, err := c.c.Do(req)
	if err != nil {
		err = fmt.Errorf("executing get minecraft mods request: %w", err)
		return &result, types.Wrap(err, "failed to execute request", -1)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var mv types.ModsResponse
	if err := parseResponse(res, "get game versions", c.opt.debug, &mv); err != nil {
		return nil, fmt.Errorf("parsing get minecraft mods response: %w", err)
	}

	return &mv, nil
	//return nil, nil
}

func (c *curseClient) buildRequestPath(path string, q ApiQueryParams) string {
	return fmt.Sprintf("%s?%s", path, q.QueryString())
}

func parseResponse[T types.CurseforgeAPIResponse](res *http.Response, step string, debug bool, out T) error {
	if res.StatusCode/100 != 2 {
		b, _ := io.ReadAll(res.Body)
		return types.Wrap(
			fmt.Errorf("%s request failed with status code %d", step, res.StatusCode),
			string(b),
			res.StatusCode,
		)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading get game versions response: %w", err)
	}

	bodyAsString := string(b)
	if debug {
		log.Printf("%s.Response (url: %s): %s", step, res.Request.URL.String(), bodyAsString)
	}

	if err := json.Unmarshal(b, out); err != nil {
		return fmt.Errorf("decoding get game versions response: %w", err)
	}
	out.SetRawResponseBody(bodyAsString)

	return nil
}
