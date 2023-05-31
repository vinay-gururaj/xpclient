package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateQueueAdapter - Create QueueAdapter
func (c *Client) CreateQueueAdapter(ctx context.Context, queueAdapter QueueAdapter) (queueAdapterId string, err error) {
	rb, err := json.Marshal(queueAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/queueadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := queueAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateQueueAdapter - Updates QueueAdapter
func (c *Client) UpdateQueueAdapter(ctx context.Context, queueAdapter QueueAdapter) (queueAdapterId string, err error) {
	rb, err := json.Marshal(queueAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/queueadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := queueAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteQueueAdapter - Deletes QueueAdapter
func (c *Client) DeleteQueueAdapter(ctx context.Context, queueAdapter QueueAdapter) (queueAdapterId string, err error) {
	rb, err := json.Marshal(queueAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/queueadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := queueAdapterId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetQueueAdapter - Fetches QueueAdapter details
func (c *Client) GetQueueAdapter(ctx context.Context, queueAdapter QueueAdapter) (queueAdapterId string, err error) {
	rb, err := json.Marshal(queueAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/queueadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := queueAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
