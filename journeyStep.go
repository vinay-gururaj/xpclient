package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateJourneySteps - Create new Journey Steps
func (c *Client) CreateJourneySteps(ctx context.Context, journeystep JourneyStep) (journeyStepId string, err error) {
	rb, err := json.Marshal(journeystep)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/journeysteps", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := journeyStepId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateJourneySteps - Updates Journey Steps
func (c *Client) UpdateJourneySteps(ctx context.Context, journeystep JourneyStep) (journeyStepId string, err error) {
	rb, err := json.Marshal(journeystep)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/journeysteps", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := journeyStepId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteJourneySteps - Deletes Journey Steps on Kitewheel
func (c *Client) DeleteJourneySteps(ctx context.Context, journeystep JourneyStep) (journeyStepId string, err error) {
	rb, err := json.Marshal(journeystep)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/journeysteps", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := journeyStepId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}
