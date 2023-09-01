package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateIdentifierType - Create new Identifier Type
func (c *Client) CreateIdentifierType(ctx context.Context, identifierType IdentifierType) (identifierTypeId string, err error) {
	rb, err := json.Marshal(identifierType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/identifiertype", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := identifierTypeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteIdentifierType - Delete Identifier Type
func (c *Client) DeleteIdentifierType(ctx context.Context, identifierType IdentifierType) (identifierTypeId string, err error) {
	rb, err := json.Marshal(identifierType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/identifiertype", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := identifierTypeId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetIdentifierType - Fetches Identifier Type within a project
func (c *Client) GetIdentifierType(ctx context.Context, IdentifierType IdentifierType) (IdentifierTypeId string, err error) {
	rb, err := json.Marshal(IdentifierType)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/identifiertype", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := IdentifierTypeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
