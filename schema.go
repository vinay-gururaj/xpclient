package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateSchema - Create Schema
func (c *Client) CreateSchema(ctx context.Context, schema ProjectSchema) (projectId string, err error) {
	rb, err := json.Marshal(schema)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/projectschema", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := projectId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateSchema - Updates Schema
func (c *Client) UpdateSchema(ctx context.Context, schema ProjectSchema) (projectId string, err error) {
	rb, err := json.Marshal(schema)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/projectschema", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := projectId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteSchema - Deletes Schema
func (c *Client) DeleteSchema(ctx context.Context, schema ProjectSchema) (projectId string, err error) {
	rb, err := json.Marshal(schema)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/projectschema", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := projectId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetSchema - Fetches Schema details
func (c *Client) GetSchema(ctx context.Context, schema ProjectSchema) (projectId string, err error) {
	rb, err := json.Marshal(schema)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projectschema", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := projectId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
