package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateKafkaListener - Create new Kafka Connection
func (c *Client) CreateKafkaListener(ctx context.Context, kafkaListener KafkaListener) (kafkaListenerResponse KafkaListener, err error) {
	rb, err := json.Marshal(kafkaListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/kafkalistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := kafkaListenerResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateKafkaListener - Updates Kafka Connection parameters
func (c *Client) UpdateKafkaListener(ctx context.Context, kafkaListener KafkaListener) (listenerId string, err error) {
	rb, err := json.Marshal(kafkaListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/kafkalistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createKafkaListener method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	println(body)

	returnedlistenerId := listenerId
	err = json.Unmarshal(body, &returnedlistenerId)
	if err != nil {
		return "", err
	}

	return returnedlistenerId, nil
}

// DeleteKafkaListener - Deletes Kafka Connection on Kitewheel
func (c *Client) DeleteKafkaListener(ctx context.Context, kafkaListener KafkaListener) (listenerId string, err error) {
	rb, err := json.Marshal(kafkaListener)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/kafkalistener", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	returnedlistenerId := listenerId
	err = json.Unmarshal(body, &returnedlistenerId)
	if err != nil {
		return "", err
	}

	return returnedlistenerId, nil
}
