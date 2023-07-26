package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateBatchConnection - Create new Batch Connection
func (c *Client) CreateBatchConnection(ctx context.Context, batchConnection BatchConnection) (batchConnectionResponse BatchConnection, err error) {
	rb, err := json.Marshal(batchConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/batchconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := batchConnectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateBatchConnection - Updates Batch Connection parameters
func (c *Client) UpdateBatchConnection(ctx context.Context, batchConnection BatchConnection) (connectorId string, err error) {
	rb, err := json.Marshal(batchConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/batchconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createBatchConnection method does not return a response
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

// DeleteBatchConnection - Deletes Batch Connection on Kitewheel
func (c *Client) DeleteBatchConnection(ctx context.Context, batchConnection BatchConnection) (connectorId string, err error) {
	rb, err := json.Marshal(batchConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/Batchconnection", c.HostURL), strings.NewReader(string(rb)))
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
