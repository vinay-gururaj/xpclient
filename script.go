package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateScript - Creates Script
func (c *Client) CreateScript(ctx context.Context, script JavaScript) (scriptId string, err error) {
	rb, err := json.Marshal(script)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/script", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := scriptId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateScript - Updates Script
func (c *Client) UpdateScript(ctx context.Context, script JavaScript) (scriptId string, err error) {
	rb, err := json.Marshal(script)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/script", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := scriptId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteScript - Deletes Script
func (c *Client) DeleteScript(ctx context.Context, script JavaScript) (scriptId string, err error) {
	rb, err := json.Marshal(script)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/script", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := scriptId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetScript - Fetches Script within a project
func (c *Client) GetScript(ctx context.Context, script JavaScript) (scriptId string, err error) {
	rb, err := json.Marshal(script)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/script", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := scriptId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
