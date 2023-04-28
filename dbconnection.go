package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateDBConnection - Create new DB Connection
func (c *Client) CreateDBConnection(ctx context.Context, dbConnection DBConnection) (dbConnectionResponse DBConnection, err error) {
	rb, err := json.Marshal(dbConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dbconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := dbConnectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateDBConnection - Updates DB Connection parameters
func (c *Client) UpdateDBConnection(ctx context.Context, dbConnection DBConnection) (connectorId string, err error) {
	rb, err := json.Marshal(dbConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/dbconnection", c.HostURL), strings.NewReader(string(rb)))
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

// DeleteDBConnection - Deletes DB Connection on Kitewheel
func (c *Client) DeleteDBConnection(ctx context.Context, dbConnection DBConnection) (connectorId string, err error) {
	rb, err := json.Marshal(dbConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/dbconnection", c.HostURL), strings.NewReader(string(rb)))
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
