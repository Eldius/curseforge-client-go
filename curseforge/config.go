package curseforge

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	apiKey  string
	baseUrl string
	timeout time.Duration
	debug   bool
}

func NewConfig(apiKey string, timeout time.Duration, debug bool) *Config {
	return &Config{
		apiKey:  apiKey,
		baseUrl: baseUrl,
		timeout: timeout,
		debug:   debug,
	}
}

func NewConfigWithBaseURL(apiKey string, baseUrl string, timeout time.Duration, debug bool) *Config {
	return &Config{
		apiKey:  apiKey,
		baseUrl: baseUrl,
		timeout: timeout,
		debug:   debug,
	}
}

func (cfg *Config) newHttpClient() *http.Client {
	return &http.Client{
		Timeout: cfg.timeout,
	}
}

func (cfg *Config) newGetRequest(c *http.Client, path string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", cfg.baseUrl, path), nil)
	if err != nil {
		log.Printf("Failed to create request object: %s", err.Error())
		return req, err
	}
	req.Header.Add(xApiKeyHeader, cfg.apiKey)

	return req, nil
}
