package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateJourneySteps - Create new Journey Steps
func (c *Client) CreateJourneyStep(ctx context.Context, fuse Fuse2) (connectionResponse KWConnections, err error) {
	rb, err := json.Marshal(fuse)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/journeystep", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := connectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateJourneySteps - Updates Journey Steps
func (c *Client) UpdateJourneyStep(ctx context.Context, connections KWConnections) (connectionResponse KWConnections, err error) {
	rb, err := json.Marshal(connections)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/journeystep", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := connectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteJourneySteps - Deletes Journey Steps on Kitewheel
func (c *Client) DeleteJourneyStep(ctx context.Context, connections KWConnections) (connectionResponse KWConnections, err error) {
	rb, err := json.Marshal(connections)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/journeystep", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := connectionResponse
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetJourneySteps - Fetches Journey Steps within a project
func (c *Client) GetJourneyStep(ctx context.Context, xpProject XPProject) (journeyStepResponse Fuse2, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/journeystep", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := journeyStepResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
