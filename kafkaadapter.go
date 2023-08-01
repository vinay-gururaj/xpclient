package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateKafkaAdapter - Create KafkaAdapter
func (c *Client) CreateKafkaAdapter(ctx context.Context, kafkaAdapter KafkaAdapter) (kafkaAdapterId string, err error) {
	rb, err := json.Marshal(kafkaAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/kafkaadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := kafkaAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateKafkaAdapter - Updates KafkaAdapter
func (c *Client) UpdateKafkaAdapter(ctx context.Context, kafkaAdapter KafkaAdapter) (kafkaAdapterId string, err error) {
	rb, err := json.Marshal(kafkaAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/kafkaadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := kafkaAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeleteKafkaAdapter - Deletes KafkaAdapter
func (c *Client) DeleteKafkaAdapter(ctx context.Context, kafkaAdapter KafkaAdapter) (kafkaAdapterId string, err error) {
	rb, err := json.Marshal(kafkaAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/kafkaadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := kafkaAdapterId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetKafkaAdapter - Fetches KafkaAdapter details
func (c *Client) GetKafkaAdapter(ctx context.Context, kafkaAdapter KafkaAdapter) (kafkaAdapterId string, err error) {
	rb, err := json.Marshal(kafkaAdapter)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/kafkaadapter", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := kafkaAdapterId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
