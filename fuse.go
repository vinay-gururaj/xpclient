package xpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// generateFuseFromProject - Generates a fuse file by reading a kitewheel project
func (c *Client) GenerateFuseFromProject(ctx context.Context, xpProject XPProject) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/fusefromproject", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	c.doRequest(req)
	if err != nil {
		panic(err)
	}
}

// generateFuseFromProject - Generates a fuse file by reading a kitewheel project
func (c *Client) GenerateProjectFromFuse(ctx context.Context, xpProject XPProject) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/projectfromfuse", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	c.doRequest(req)
	if err != nil {
		panic(err)
	}
}
