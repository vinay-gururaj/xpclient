package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateDatabaseAdapter - Create DatabaseAdapter
func (c *Client) CreateDatabaseAdapter(ctx context.Context, databaseAdapter DatabaseAdapter) (databaseAdapterId string, err error) {
	rb, err := json.Marshal(databaseAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/databaseadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := databaseAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateDatabaseAdapter - Updates DatabaseAdapter
func (c *Client) UpdateDatabaseAdapter(ctx context.Context, databaseAdapter DatabaseAdapter) (databaseAdapterId string, err error) {
	rb, err := json.Marshal(databaseAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/databaseadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := databaseAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteDatabaseAdapter - Deletes DatabaseAdapter
func (c *Client) DeleteDatabaseAdapter(ctx context.Context, databaseAdapter DatabaseAdapter) (databaseAdapterId string, err error) {
	rb, err := json.Marshal(databaseAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/databaseadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := databaseAdapterId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetDatabaseAdapter - Fetches DatabaseAdapter details
func (c *Client) GetDatabaseAdapter(ctx context.Context, databaseAdapter DatabaseAdapter) (databaseAdapterId string, err error) {
	rb, err := json.Marshal(databaseAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/databaseadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := databaseAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
