package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateEnvironment - Create new Environment on Kitewheel
func (c *Client) CreateEnvironment(ctx context.Context, environment Environment) (environmentId string, err error) {
	rb, err := json.Marshal(environment)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/environment", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return "", err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	returnedEnvironment := environmentId
	err = json.Unmarshal(body, &returnedEnvironment)
	if err != nil {
		return "", err
	}

	return returnedEnvironment, nil
}

// UpdateXPProject - Updates Project Name on Kitewheel
func (c *Client) UpdateEnvironment(ctx context.Context, environment Environment) (environmentId string, err error) {
	rb, err := json.Marshal(environment)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/environment", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	returnedEnvironment := environmentId
	err = json.Unmarshal(body, &returnedEnvironment)
	if err != nil {
		return "", err
	}

	return returnedEnvironment, nil

}

// DeleteXPProject - Deletes Project on Kitewheel
func (c *Client) DeleteEnvironment(ctx context.Context, environment Environment) (environmentId string, err error) {
	rb, err := json.Marshal(environment)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/environment", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	returnedEnvironment := environmentId
	err = json.Unmarshal(body, &returnedEnvironment)
	if err != nil {
		return "", err
	}

	return returnedEnvironment, nil
}
