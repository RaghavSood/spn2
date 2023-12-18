package spn2

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://web.archive.org/save"

// Client represents a client to interact with the SPN2 API.
type Client struct {
	httpClient *http.Client
	AccessKey  string
	SecretKey  string
}

// NewClient creates a new SPN2 API client.
func NewClient(accessKey, secretKey string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		AccessKey:  accessKey,
		SecretKey:  secretKey,
	}
}

// SubmitURL submits a new URL to the SPN2 API for capturing.
func (c *Client) SubmitURL(url string) (*CaptureResponse, error) {
	reqBody := bytes.NewBufferString("url=" + url)
	req, err := http.NewRequest("POST", baseURL, reqBody)
	if err != nil {
		return nil, err
	}

	c.addAuthHeaders(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var captureResp CaptureResponse
	if err := json.NewDecoder(resp.Body).Decode(&captureResp); err != nil {
		return nil, err
	}

	return &captureResp, nil
}

// GetStatus retrieves the status of a capture using the provided Job ID.
func (c *Client) GetStatus(jobID string) (*StatusResponse, error) {
	reqURL := fmt.Sprintf("%s/status/%s", baseURL, jobID)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	c.addAuthHeaders(req)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var statusResp StatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, err
	}

	return &statusResp, nil
}

// GetSystemStatus checks the system status of the SPN2 API.
func (c *Client) GetSystemStatus() (*SystemStatusResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/status/system", baseURL), nil)
	if err != nil {
		return nil, err
	}

	c.addAuthHeaders(req)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get system status")
	}

	var systemStatus SystemStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&systemStatus); err != nil {
		return nil, err
	}

	return &systemStatus, nil
}

// GetUserStatus checks the user's account status in the SPN2 API.
func (c *Client) GetUserStatus() (*UserStatusResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/status/user", baseURL), nil)
	if err != nil {
		return nil, err
	}

	c.addAuthHeaders(req)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get user status")
	}

	var userStatus UserStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&userStatus); err != nil {
		return nil, err
	}

	return &userStatus, nil
}

// addAuthHeaders adds the necessary authentication headers to the request.
func (c *Client) addAuthHeaders(req *http.Request) {
	authValue := fmt.Sprintf("LOW %s:%s", c.AccessKey, c.SecretKey)
	req.Header.Set("Authorization", authValue)
}
