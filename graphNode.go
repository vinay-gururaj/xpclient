package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateGraphNode - Create GraphNode
func (c *Client) CreateGraphNode(ctx context.Context, graphNode SimpleNode) (graphNodeId string, err error) {
	rb, err := json.Marshal(graphNode)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/graphNode", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := graphNodeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateGraphNode - Updates GraphNode
func (c *Client) UpdateGraphNode(ctx context.Context, graphNode SimpleNode) (graphNodeId string, err error) {
	rb, err := json.Marshal(graphNode)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/graphNode", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := graphNodeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteGraphNode - Deletes GraphNode
func (c *Client) DeleteGraphNode(ctx context.Context, graphNode SimpleNode) (graphNodeId string, err error) {
	rb, err := json.Marshal(graphNode)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/graphNode", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := graphNodeId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetGraphNode - Fetches Graph Node details
func (c *Client) GetGraphNode(ctx context.Context, graphNode SimpleNode) (graphNodeId string, err error) {
	rb, err := json.Marshal(graphNode)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/graphNode", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := graphNodeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
