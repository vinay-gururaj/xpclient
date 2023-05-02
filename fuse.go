package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// generateFuseFromProject - Generates a fuse file by reading a kitewheel project
func (c *Client) GenerateFuseFromProject(ctx context.Context, xpProject XPProject) (status bool, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/fuse", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return false, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	responseStatus := status
	err = json.Unmarshal(body, &responseStatus)
	if err != nil {
		return false, err
	}

	return responseStatus, nil
}

// generateFuseFromProject - Generates a fuse file by reading a kitewheel project
func (c *Client) GenerateProjectFromFuse(ctx context.Context, xpProject XPProject) (status bool, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/fuse", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return false, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	responseStatus := status
	err = json.Unmarshal(body, &responseStatus)
	if err != nil {
		return false, err
	}

	return responseStatus, nil
}
