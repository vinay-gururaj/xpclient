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

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/project", c.HostURL), strings.NewReader(string(rb)))
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

// CreateXPProject - Create new CreateXPProject
func (c *Client) GetProject(ctx context.Context, xpProject XPProject) (xpProjectResponse KWProject, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/project", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	xpProjectResp := xpProjectResponse
	err = json.Unmarshal(body, &xpProjectResp)
	if err != nil {
		panic(err)
	}

	return xpProjectResp, nil
}

// UpdateXPProject - Updates Project Name on Kitewheel
func (c *Client) UpdateXPProject(ctx context.Context, xpProject XPProject) (projectId string, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/project", c.HostURL), strings.NewReader(string(rb)))
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

// DeleteXPProject - Deletes Project on Kitewheel
func (c *Client) DeleteXPProject(ctx context.Context, xpProject XPProject) (projectId string, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/project", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
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
