package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateS3Connection - Create new s3 Connection
func (c *Client) CreateS3Connection(ctx context.Context, s3Connection S3Connection) (s3ConnectionResponse S3Connection, err error) {
	rb, err := json.Marshal(s3Connection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/s3connection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := s3ConnectionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateS3Connection - Updates s3 Connection parameters
func (c *Client) UpdateS3Connection(ctx context.Context, s3Connection S3Connection) {
	rb, err := json.Marshal(s3Connection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/s3connection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createS3Connection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	println(body)
}

// DeleteS3Connection - Deletes s3 Connection on Kitewheel
func (c *Client) DeleteS3Connection(ctx context.Context, s3Connection S3Connection) (connectionId string, err error) {
	rb, err := json.Marshal(s3Connection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/s3connection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := connectionId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}
