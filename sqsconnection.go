package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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
