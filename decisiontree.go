package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreatedecisionTree - Creates decisionTree
func (c *Client) CreateDecisionTree(ctx context.Context, decisionTree DecisionTree) (decisionTreeId string, err error) {
	rb, err := json.Marshal(decisionTree)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/decisiontree", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := decisionTreeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdatedecisionTree - Updates decisionTree
func (c *Client) UpdateDecisionTree(ctx context.Context, decisionTree DecisionTree) (decisionTreeId string, err error) {
	rb, err := json.Marshal(decisionTree)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/decisiontree", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := decisionTreeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeletedecisionTree - Deletes decisionTree
func (c *Client) DeleteDecisionTree(ctx context.Context, decisionTree DecisionTree) (decisionTreeId string, err error) {
	rb, err := json.Marshal(decisionTree)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/decisiontree", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := decisionTreeId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}

// GetdecisionTree - Fetches decisionTree details
func (c *Client) GetDecisionTree(ctx context.Context, decisionTree DecisionTree) (decisionTreeId string, err error) {
	rb, err := json.Marshal(decisionTree)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/decisiontree", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := decisionTreeId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
