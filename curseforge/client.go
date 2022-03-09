package curseforge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	defaultTimeout = 30 * time.Second
	baseUrl        = "https://api.curseforge.com"
	xApiKeyHeader  = "x-api-key"

	modSearchPath      = "/v1/mods/search"
	gamesListPath      = "/v1/games"
	categoriesListPath = "/v1/categories"
)

type Client struct {
	cfg *Config
}

func NewClientWithConfig(cfg *Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

func NewClient(apiKey string) *Client {
	return NewClientWithConfig(&Config{
		apiKey:  apiKey,
		timeout: defaultTimeout,
		baseUrl: baseUrl,
	})
}

func (_c *Client) GetGames() (GamesResponse, error) {
	var result GamesResponse
	c := _c.cfg.newHttpClient()

	req, err := _c.cfg.newGetRequest(c, gamesListPath)
	if err != nil {
		log.Printf("Failed to create request object: %s", err.Error())
		return result, err
	}

	res, err := c.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %s", err.Error())
		return result, err
	}
	_c.parseResponse(res, &result)

	return result, nil
}

func (_c *Client) GetMods(gameId string) (ModsResponse, error) {
	var result ModsResponse
	c := _c.cfg.newHttpClient()

	req, err := _c.cfg.newGetRequest(c, modSearchPath)
	if err != nil {
		log.Printf("Failed to create request object: %s", err.Error())
		return result, err
	}

	q := req.URL.Query()
	q.Add("gameId", gameId)
	req.URL.RawQuery = q.Encode()
	log.Println(req.URL.String())
	res, err := c.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %s", err.Error())
		return result, err
	}
	_c.parseResponse(res, &result)

	log.Println("code:", res.StatusCode)

	return result, nil
}

func (_c *Client) GetModsByCategory(gameId string, modCategorySlug string, searchFilter string) (*ModsResponse, error) {
	var result *ModsResponse
	c := _c.cfg.newHttpClient()

	req, err := _c.cfg.newGetRequest(c, modSearchPath)
	if err != nil {
		log.Printf("Failed to create request object: %s", err.Error())
		return result, err
	}

	q := req.URL.Query()
	q.Add("gameId", gameId)
	q.Add("categoryId", modCategorySlug)
	q.Add("searchFilter", searchFilter)
	req.URL.RawQuery = q.Encode()
	log.Println(req.URL.String())
	res, err := c.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %s", err.Error())
		return result, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http response error: " + res.Status)
	}

	_c.parseResponse(res, &result)

	return result, nil
}

func (_c *Client) GetCategories(gameId string) (*ModsResponse, error) {
	var result *ModsResponse

	c := _c.cfg.newHttpClient()

	req, err := _c.cfg.newGetRequest(c, categoriesListPath)
	if err != nil {
		log.Printf("Failed to create request object: %s", err.Error())
		return result, err
	}

	q := req.URL.Query()
	q.Add("gameId", gameId)
	req.URL.RawQuery = q.Encode()
	log.Println(req.URL.String())
	res, err := c.Do(req)
	if err != nil {
		log.Printf("Failed to execute request: %s", err.Error())
		return result, err
	}
	_c.parseResponse(res, &result)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http response error: " + res.Status)
	}

	return result, nil
}

func (c *Client) debugResponse(r *http.Response) {
	if c.cfg.debug {
		reader := r.Body
		defer func() {
			err := reader.Close()
			if err != nil {
				log.Println("Failed to close reader:", err.Error())
			}
		}()
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Printf("Failed to execute request: %s", err.Error())
			return
		}
		msg := "---\nheaders:\n"
		for k, v := range r.Header {
			msg += fmt.Sprintf(" - %s: [%s]\n", k, strings.Join(v, ", "))
		}
		msg += fmt.Sprintf("response:\n%s\n---", string(b))
		log.Println(msg)
		log.Println()
		r.Body = io.NopCloser(bytes.NewReader(b))
	}
}

func (c *Client) parseResponse(r *http.Response, result interface{}) {
	c.debugResponse(r)
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(result)
}
