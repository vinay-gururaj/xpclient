package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateDatabaseListener - Create DatabaseListener
func (c *Client) CreateDatabaseListener(ctx context.Context, databaseListener DatabaseListener) (databaseListenerId string, err error) {
	rb, err := json.Marshal(databaseListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/databaselistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := databaseListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateDatabaseListener - Updates DatabaseListener
func (c *Client) UpdateDatabaseListener(ctx context.Context, databaseListener DatabaseListener) (databaseListenerId string, err error) {
	rb, err := json.Marshal(databaseListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/databaselistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := databaseListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteDatabaseListener - Deletes DatabaseListener
func (c *Client) DeleteDatabaseListener(ctx context.Context, databaseListener DatabaseListener) (databaseListenerId string, err error) {
	rb, err := json.Marshal(databaseListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/databaselistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := databaseListenerId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetDatabaseListener - Fetches DatabaseListener details
func (c *Client) GetDatabaseListener(ctx context.Context, databaseListener DatabaseListener) (databaseListenerId string, err error) {
	rb, err := json.Marshal(databaseListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/databaselistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := databaseListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
