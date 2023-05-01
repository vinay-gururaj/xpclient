package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateScripts - Creates Scripts
func (c *Client) CreateScripts(ctx context.Context, fuse Fuse2) (connectionResponse KWConnections, err error) {
	rb, err := json.Marshal(fuse)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/scripts", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := connectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateScripts - Updates Scripts
func (c *Client) UpdateScripts(ctx context.Context, connections KWConnections) (connectionResponse KWConnections, err error) {
	rb, err := json.Marshal(connections)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/scripts", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := connectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteScripts - Deletes Scripts
func (c *Client) DeleteScripts(ctx context.Context, connections KWConnections) (connectionResponse KWConnections, err error) {
	rb, err := json.Marshal(connections)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/scripts", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := connectionResponse
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetScripts - Fetches Scripts within a project
func (c *Client) GetScripts(ctx context.Context, xpProject XPProject) (graphResponse Fuse2, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/scripts", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := graphResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
