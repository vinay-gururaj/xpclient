package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateBatchListener - Create BatchListener
func (c *Client) CreateBatchListener(ctx context.Context, batchListener BatchListener) (batchListenerId string, err error) {
	rb, err := json.Marshal(batchListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/batchlistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := batchListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateBatchListener - Updates BatchListener
func (c *Client) UpdateBatchListener(ctx context.Context, batchListener BatchListener) (batchListenerId string, err error) {
	rb, err := json.Marshal(batchListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/batchlistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := batchListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteBatchListener - Deletes BatchListener
func (c *Client) DeleteBatchListener(ctx context.Context, batchListener BatchListener) (batchListenerId string, err error) {
	rb, err := json.Marshal(batchListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/batchlistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := batchListenerId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetBatchListener - Fetches BatchListener details
func (c *Client) GetBatchListener(ctx context.Context, batchListener BatchListener) (batchListenerId string, err error) {
	rb, err := json.Marshal(batchListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/batchlistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := batchListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
