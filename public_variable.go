package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreatePublicVariable - Create new CreatePublicVariable
func (c *Client) CreatePublicVariable(ctx context.Context, publicVariable PublicVariable) (publicVariableId string, err error) {
	rb, err := json.Marshal(publicVariable)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/publicvariable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return "", err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	respPublicVariableId := publicVariableId
	err = json.Unmarshal(body, &respPublicVariableId)
	if err != nil {
		return "", err
	}

	return respPublicVariableId, nil
}

// CreatePublicVariable - Create new CreatePublicVariable
func (c *Client) GetPublicVariable(ctx context.Context, publicVariable PublicVariable) (publicVariableResponse PublicVariable, err error) {
	rb, err := json.Marshal(publicVariable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/publicvariable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	respPublicVariableId := publicVariableResponse
	err = json.Unmarshal(body, &respPublicVariableId)
	if err != nil {
		panic(err)
	}

	return respPublicVariableId, nil
}

// UpdatePublicVariable - Updates Public Variable Name on Kitewheel
func (c *Client) UpdatePublicVariable(ctx context.Context, publicVariable PublicVariable) (publicVariableId string, err error) {
	rb, err := json.Marshal(publicVariable)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/publicvariable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return "", err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	respPublicVariableId := publicVariableId
	err = json.Unmarshal(body, &respPublicVariableId)
	if err != nil {
		return "", err
	}

	return respPublicVariableId, nil

}

// DeletePublicVariable - Deletes Public Variable on Kitewheel
func (c *Client) DeletePublicVariable(ctx context.Context, publicVariable PublicVariable) (publicVariableId string, err error) {
	rb, err := json.Marshal(publicVariable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/publicvariable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	respPublicVariableId := publicVariableId
	err = json.Unmarshal(body, &respPublicVariableId)
	if err != nil {
		return "", err
	}

	return respPublicVariableId, nil

}
