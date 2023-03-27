package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateXPProject - Create new CreateXPProject
func (c *Client) CreateXPProject(ctx context.Context, xpProject XPProject) (projectId string, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/createXPProject", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return "", err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	xpProjectId := projectId
	err = json.Unmarshal(body, &xpProjectId)
	if err != nil {
		return "", err
	}

	return xpProjectId, nil
}

// CreateSQSConnection - Create new SQS Connection
func (c *Client) CreateSQSConnection(ctx context.Context, sqsConnection SQSConnection) (connectionId string, err error) {
	rb, err := json.Marshal(sqsConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/createSQSConnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createSQSConnection method does not return a response
	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	xpConnectionId := connectionId
	err = json.Unmarshal(body, &xpConnectionId)
	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}
