package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateApiListener - Create ApiListener
func (c *Client) CreateApiListener(ctx context.Context, apiListener APIListener) (apiListenerId string, err error) {
	rb, err := json.Marshal(apiListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apilistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := apiListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateApiListener - Updates ApiListener
func (c *Client) UpdateApiListener(ctx context.Context, apiListener APIListener) (apiListenerId string, err error) {
	rb, err := json.Marshal(apiListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apilistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := apiListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteApiListener - Deletes ApiListener
func (c *Client) DeleteApiListener(ctx context.Context, apiListener APIListener) (apiListenerId string, err error) {
	rb, err := json.Marshal(apiListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apilistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := apiListenerId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetApiListener - Fetches ApiListener details
func (c *Client) GetApiListener(ctx context.Context, apiListener APIListener) (apiListenerId string, err error) {
	rb, err := json.Marshal(apiListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apilistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := apiListenerId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
