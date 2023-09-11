package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateEnvironment - Create new Environment on Kitewheel
func (c *Client) CreateConditional(ctx context.Context, conditional ConditionalGraph) (conditionalId string, err error) {
	rb, err := json.Marshal(conditional)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/conditional", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return "", err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	returnedConditional := conditionalId
	err = json.Unmarshal(body, &returnedConditional)
	if err != nil {
		return "", err
	}

	return returnedConditional, nil
}

// UpdateXPProject - Updates Project Name on Kitewheel
func (c *Client) UpdateConditional(ctx context.Context, conditional ConditionalGraph) (conditionalId string, err error) {
	rb, err := json.Marshal(conditional)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/conditional", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	returnedConditional := conditionalId
	err = json.Unmarshal(body, &returnedConditional)
	if err != nil {
		return "", err
	}

	return returnedConditional, nil

}

// DeleteXPProject - Deletes Project on Kitewheel
func (c *Client) DeleteConditional(ctx context.Context, conditional ConditionalGraph) (conditionalId string, err error) {
	rb, err := json.Marshal(conditional)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/conditional", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	returnedConditional := conditionalId
	err = json.Unmarshal(body, &returnedConditional)
	if err != nil {
		return "", err
	}

	return returnedConditional, nil
}
