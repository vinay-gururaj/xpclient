package xpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "http://localhost:8001"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	SessionId  string
	CsrfToken  string
}

// NewClient -
func NewClient(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	c.SessionId = "s:xmHjJr_P3-NlEk2YlaaRDZd2UAPwz6MV.CEqSnSZQzGmFRu1ZjAHI2OtEZvSTrG+p1P9VMiwss2E"
	c.CsrfToken = "nDmldoDQ-ElCaoyNPPweeTcUlhrShDW8SsfM"

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("SessionId", c.SessionId)
	req.Header.Set("CsrfToken", c.CsrfToken)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
