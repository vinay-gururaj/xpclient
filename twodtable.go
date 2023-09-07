package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateTwoDTable - Creates TwoDTable
func (c *Client) CreateTwoDTable(ctx context.Context, twoDTable TwoDTable) (twoDTableId string, err error) {
	rb, err := json.Marshal(twoDTable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/twodtable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := twoDTableId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateTwoDTable - Updates TwoDTable
func (c *Client) UpdateTwoDTable(ctx context.Context, twoDTable TwoDTable) (twoDTableId string, err error) {
	rb, err := json.Marshal(twoDTable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/twodtable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := twoDTableId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteTwoDTable - Deletes TwoDTable
func (c *Client) DeleteParentNode(ctx context.Context, twoDTable TwoDTable) (twoDTableId string, err error) {
	rb, err := json.Marshal(twoDTable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/twodtable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := twoDTableId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetTwoDTable - Fetches TwoDTable details
func (c *Client) GetTwoDTable(ctx context.Context, twoDTable TwoDTable) (twoDTableId string, err error) {
	rb, err := json.Marshal(twoDTable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/twodtable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := twoDTableId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
