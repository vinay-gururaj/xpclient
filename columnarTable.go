package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateColumnarTable - Create new ColumnarTable
func (c *Client) CreateColumnarTable(ctx context.Context, columnarTable ColumnarTable) (columnarTableId string, err error) {
	rb, err := json.Marshal(columnarTable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/columnartable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := columnarTableId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// UpdateColumnarTable - Updates ColumnarTable
func (c *Client) UpdateColumnarTable(ctx context.Context, columnarTable ColumnarTable) (columnarTableId string, err error) {
	rb, err := json.Marshal(columnarTable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/columnartable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)

	if err != nil {
		panic(err)
	}

	response := columnarTableId
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

// DeletecolumnarTable - Delete columnarTable
func (c *Client) DeleteColumnarTable(ctx context.Context, columnarTable ColumnarTable) (columnarTableId string, err error) {
	rb, err := json.Marshal(columnarTable)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/columnartable", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		panic(err)
	}

	println(body)

	xpConnectionId := columnarTableId
	err = json.Unmarshal(body, &xpConnectionId)

	if err != nil {
		panic(err)
	}
	return xpConnectionId, nil
}
