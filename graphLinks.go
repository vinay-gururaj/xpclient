package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateGraphNodeLink - Create GraphNodeLink
func (c *Client) CreateGraphNodeLink(ctx context.Context, graphNodeLink GraphLink) (graphNodeLinkId string, err error) {
	rb, err := json.Marshal(graphNodeLink)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/link", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := graphNodeLinkId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateGraphNodeLink - Updates GraphNodeLink
func (c *Client) UpdateGraphNodeLink(ctx context.Context, graphNodeLink GraphLink) (graphNodeLinkId string, err error) {
	rb, err := json.Marshal(graphNodeLink)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/link", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := graphNodeLinkId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteGraphNodeLink - Deletes GraphNodeLink
func (c *Client) DeleteGraphNodeLink(ctx context.Context, graphNodeLink GraphLink) (graphNodeLinkId string, err error) {
	rb, err := json.Marshal(graphNodeLink)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/link", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := graphNodeLinkId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetGraphNodeLink - Fetches Graph Node Link details
func (c *Client) GetGraphNodeLink(ctx context.Context, graphNodeLink GraphLink) (graphNodeLinkId string, err error) {
	rb, err := json.Marshal(graphNodeLink)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/link", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := graphNodeLinkId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
