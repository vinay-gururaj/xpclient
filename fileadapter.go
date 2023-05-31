package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateFileAdapter - Create FileAdapter
func (c *Client) CreateFileAdapter(ctx context.Context, fileAdapter FileAdapter) (fileAdapterId string, err error) {
	rb, err := json.Marshal(fileAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/fileadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := fileAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateFileAdapter - Updates FileAdapter
func (c *Client) UpdateFileAdapter(ctx context.Context, fileAdapter FileAdapter) (fileAdapterId string, err error) {
	rb, err := json.Marshal(fileAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/fileadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := fileAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteFileAdapter - Deletes FileAdapter
func (c *Client) DeleteFileAdapter(ctx context.Context, fileAdapter FileAdapter) (fileAdapterId string, err error) {
	rb, err := json.Marshal(fileAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/fileadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := fileAdapterId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetFileAdapter - Fetches FileAdapter details
func (c *Client) GetFileAdapter(ctx context.Context, fileAdapter FileAdapter) (fileAdapterId string, err error) {
	rb, err := json.Marshal(fileAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fileadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := fileAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
