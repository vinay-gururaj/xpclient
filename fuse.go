package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// generateFuseFromProject - Generates a fuse file by reading a kitewheel project
func (c *Client) generateFuseFromProject(ctx context.Context, xpProject XPProject) (status bool, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/fuse", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return false, err
	}

	body, err := c.doRequest(req)

	err = json.Unmarshal(body, nil)
	if err != nil {
		return false, err
	}

	return true, nil
}

// generateFuseFromProject - Generates a fuse file by reading a kitewheel project
func (c *Client) generateProjectFromFuse(ctx context.Context, xpProject XPProject) (status bool, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/fuse", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return false, err
	}

	body, err := c.doRequest(req)

	err = json.Unmarshal(body, nil)
	if err != nil {
		return false, err
	}

	return true, nil
}
