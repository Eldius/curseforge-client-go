package config

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	// CurseforgeBaseURL default API base path
	CurseforgeBaseURL = "https://api.curseforge.com"

	defaultTimeout = 30 * time.Second
	xAPIKeyHeader  = "x-api-key"
)

// Config is the client config data
type Config struct {
	apiKey  string
	baseURL string
	timeout time.Duration
	debug   bool
}

// NewDefaultConfig creates a new client Config with default values
func NewDefaultConfig(apiKey string) *Config {
	return NewConfigWithBaseURL(
		apiKey,
		CurseforgeBaseURL,
		defaultTimeout,
		false,
	)
}

// NewConfig creates a new client Config
func NewConfig(apiKey string, timeout time.Duration, debug bool) *Config {
	return NewConfigWithBaseURL(
		apiKey,
		CurseforgeBaseURL,
		timeout,
		debug,
	)
}

// NewConfigWithBaseURL creates a new client Config
func NewConfigWithBaseURL(apiKey string, baseURL string, timeout time.Duration, debug bool) *Config {
	return &Config{
		apiKey:  apiKey,
		baseURL: baseURL,
		timeout: timeout,
		debug:   debug,
	}
}

// IsDebug is debug enabled
func (cfg *Config) IsDebug() bool {
	return cfg.debug
}

// NewHTTPClient creates a new HTTP client
func (cfg *Config) NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: cfg.timeout,
	}
}

// NewGetRequest creates a new GET request object to be used with client
func (cfg *Config) NewGetRequest(c *http.Client, path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", cfg.baseURL, path), nil)
	if err != nil {
		log.Printf("Failed to create request object: %s", err.Error())
		return req, err
	}
	req.Header.Add(xAPIKeyHeader, cfg.apiKey)

	return req, nil
}
