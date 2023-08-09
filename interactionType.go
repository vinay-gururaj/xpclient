package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateInteractionType - Create new Interaction Type
func (c *Client) CreateInteractionType(ctx context.Context, interactionType InteractionType) (interactionTypeId string, err error) {
	rb, err := json.Marshal(interactionType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/interactiontype", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := interactionTypeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateInteractionType - Updates Interaction Type
func (c *Client) UpdateInteractionType(ctx context.Context, interactionType InteractionType) (interactionTypeId string, err error) {
	rb, err := json.Marshal(interactionType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/interactiontype", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := interactionTypeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteInteractionType - Delete Interaction Type
func (c *Client) DeleteInteractionType(ctx context.Context, interactionType InteractionType) (interactionTypeId string, err error) {
	rb, err := json.Marshal(interactionType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/interactiontype", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := interactionTypeId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetInteractionType - Fetches Interaction Type within a project
func (c *Client) GetInteractionType(ctx context.Context, interactionType InteractionType) (interactionTypeId string, err error) {
	rb, err := json.Marshal(interactionType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/interactiontype", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := interactionTypeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
