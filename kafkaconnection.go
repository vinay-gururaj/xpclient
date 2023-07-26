package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateKafkaConnection - Create new Kafka Connection
func (c *Client) CreateKafkaConnection(ctx context.Context, kafkaConnection KafkaConnection) (kafkaConnectionResponse KafkaConnection, err error) {
	rb, err := json.Marshal(kafkaConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/kafkaconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := kafkaConnectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateKafkaConnection - Updates Kafka Connection parameters
func (c *Client) UpdateKafkaConnection(ctx context.Context, kafkaConnection KafkaConnection) (connectorId string, err error) {
	rb, err := json.Marshal(kafkaConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/kafkaconnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createKafkaConnection method does not return a response
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

// DeleteKafkaConnection - Deletes Kafka Connection on Kitewheel
func (c *Client) DeleteKafkaConnection(ctx context.Context, kafkaConnection KafkaConnection) (connectorId string, err error) {
	rb, err := json.Marshal(kafkaConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/kafkaconnection", c.HostURL), strings.NewReader(string(rb)))
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
