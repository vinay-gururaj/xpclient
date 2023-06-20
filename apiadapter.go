package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateApiAdapter - Create ApiAdapter
func (c *Client) CreateApiAdapter(ctx context.Context, apiAdapter WebServiceAdapter) (apiAdapterId string, err error) {
	rb, err := json.Marshal(apiAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apiadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := apiAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateApiAdapter - Updates ApiAdapter
func (c *Client) UpdateApiAdapter(ctx context.Context, apiAdapter WebServiceAdapter) (apiAdapterId string, err error) {
	rb, err := json.Marshal(apiAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apiadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := apiAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteApiAdapter - Deletes ApiAdapter
func (c *Client) DeleteApiAdapter(ctx context.Context, apiAdapter WebServiceAdapter) (apiAdapterId string, err error) {
	rb, err := json.Marshal(apiAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apiadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := apiAdapterId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetApiAdapter - Fetches ApiAdapter details
func (c *Client) GetApiAdapter(ctx context.Context, apiAdapter WebServiceAdapter) (apiAdapterId string, err error) {
	rb, err := json.Marshal(apiAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apiadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := apiAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
