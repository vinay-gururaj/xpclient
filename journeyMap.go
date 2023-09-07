package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateJourneyMap - Create Journey Map
func (c *Client) CreateJourneyMap(ctx context.Context, journeyMap JourneyMap) (journeyMapId string, err error) {
	rb, err := json.Marshal(journeyMap)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/journeymap", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := journeyMapId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateJourneyMap - Updates Journey Map
func (c *Client) UpdateJourneyMap(ctx context.Context, journeyMap JourneyMap) (journeyMapId string, err error) {
	rb, err := json.Marshal(journeyMap)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/journeymap", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := journeyMapId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteJourneyMap - Deletes Journey Map
func (c *Client) DeleteJourneyMap(ctx context.Context, journeyMap JourneyMap) (journeyMapId string, err error) {
	rb, err := json.Marshal(journeyMap)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/journeymap", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := journeyMapId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetJourneyMap - Fetches Journey Map details
func (c *Client) GetJourneyMap(ctx context.Context, journeyMap JourneyMap) (journeyMapId string, err error) {
	rb, err := json.Marshal(journeyMap)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/journeymap", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := journeyMapId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
