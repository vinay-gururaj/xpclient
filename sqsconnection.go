package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateSQSConnection - Create new SQS Connection
func (c *Client) CreateSQSConnection(ctx context.Context, sqsConnection SQSConnection) (sqsConnectionResponse SQSConnection, err error) {
	rb, err := json.Marshal(sqsConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/sqsconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := sqsConnectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateSQSConnection - Updates SQS Connection parameters
func (c *Client) UpdateSQSConnection(ctx context.Context, sqsConnection SQSConnection) (connectorId string, err error) {
	rb, err := json.Marshal(sqsConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/sqsconnection", c.HostURL), strings.NewReader(string(rb)))
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

// DeleteSQSConnection - Deletes SQS Connection on Kitewheel
func (c *Client) DeleteSQSConnection(ctx context.Context, sqsConnection SQSConnection) (connectorId string, err error) {
	rb, err := json.Marshal(sqsConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/sqsconnection", c.HostURL), strings.NewReader(string(rb)))
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
