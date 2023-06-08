package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateAPIConnection - Create new API Connection
func (c *Client) CreateAPIConnection(ctx context.Context, apiConnection APIConnection) (apiConnectionResponse APIConnection, err error) {
	rb, err := json.Marshal(apiConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/apiconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := apiConnectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateAPIConnection - Updates API Connection parameters
func (c *Client) UpdateAPIConnection(ctx context.Context, apiConnection APIConnection) (connectorId string, err error) {
	rb, err := json.Marshal(apiConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/apiconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	returnedConnectorId := connectorId
	err = json.Unmarshal(body, &returnedConnectorId)
	if err != nil {
		return "", err
	}

	return returnedConnectorId, nil

}

// DeleteAPIConnection - Deletes API Connection on Kitewheel
func (c *Client) DeleteAPIConnection(ctx context.Context, apiConnection APIConnection) (connectorId string, err error) {
	rb, err := json.Marshal(apiConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/apiconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	returnedConnectorId := connectorId
	err = json.Unmarshal(body, &returnedConnectorId)
	if err != nil {
		return "", err
	}

	return returnedConnectorId, nil
}
