package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIClient represents the client to make API requests
type APIClient struct {
	baseURL string
	client  *http.Client
}

// NewAPIClient creates a new APIClient with the specified base URL
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// Get performs a GET request to the API with the specified endpoint
func (c *APIClient) Get(endpoint string, result interface{}) error {
	url := c.baseURL + endpoint

	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}

	return nil
}
