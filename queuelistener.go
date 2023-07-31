package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateQueueListener - Create QueueListener
func (c *Client) CreateQueueListener(ctx context.Context, queueListener QueueListener) (queueListenerId string, err error) {
	rb, err := json.Marshal(queueListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/queuelistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := queueListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateQueueListener - Updates QueueListener
func (c *Client) UpdateQueueListener(ctx context.Context, queueListener QueueListener) (queueListenerId string, err error) {
	rb, err := json.Marshal(queueListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/queuelistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := queueListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteQueueListener - Deletes QueueListener
func (c *Client) DeleteQueueListener(ctx context.Context, queueListener QueueListener) (queueListenerId string, err error) {
	rb, err := json.Marshal(queueListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/queuelistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := queueListenerId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetQueueListener - Fetches QueueListener details
func (c *Client) GetQueueListener(ctx context.Context, queueListener QueueListener) (queueListenerId string, err error) {
	rb, err := json.Marshal(queueListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/queuelistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := queueListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
