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
