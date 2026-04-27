package serpentapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.apiserpent.com/search"

// Client is an HTTP client for the Serpent API.
type Client struct {
	APIKey     string
	HTTPClient *http.Client
	Engine     string
}

// NewClient creates a new Serpent API client.
// apiKey is required. engine defaults to "google" if empty.
func NewClient(apiKey, engine string) *Client {
	if engine == "" {
		engine = "google"
	}
	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		Engine: engine,
	}
}

// Search performs a search query and returns structured results.
func (c *Client) Search(query string) (*SearchResponse, error) {
	if c.APIKey == "" {
		return nil, fmt.Errorf("serpentapi: API key is required")
	}
	if query == "" {
		return nil, fmt.Errorf("serpentapi: query must not be empty")
	}

	params := url.Values{}
	params.Set("engine", c.Engine)
	params.Set("q", query)
	params.Set("api_key", c.APIKey)

	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("serpentapi: failed to create request: %w", err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("serpentapi: request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("serpentapi: failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("serpentapi: API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result SearchResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("serpentapi: failed to parse response: %w", err)
	}

	if result.Error != "" {
		return nil, fmt.Errorf("serpentapi: API error: %s", result.Error)
	}

	return &result, nil
}

// SearchSpiderWebs is a convenience method that searches for different types
// of spider webs. This is useful for comparing web structures across species.
func (c *Client) SearchSpiderWebs(webType string) (*SearchResponse, error) {
	query := fmt.Sprintf("%s spider web structure geometry", webType)
	return c.Search(query)
}
