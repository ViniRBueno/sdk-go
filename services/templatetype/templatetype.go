package template

import (
	"bytes"
	"encoding/json"
	"fmt"

	common "github.com/ViniRBueno/sdk-go/services/common"
)

// Client is the client for Tenant Discovery Rest Api
type Client struct {
	common.BaseClient
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, token string) *Client {

	return &Client{
		common.NewBase(
			hostname,
			token,
		)}
}

// Get gets all the templatetypes
func (c *Client) Get() (*TemplateTypesList, error) {

	body, err := c.HTTPRequest(fmt.Sprintf("/v1/campaigns/catalogTemplateTypes"), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	templatetypes := &TemplateTypesList{}
	err = json.NewDecoder(body).Decode(templatetypes)

	if err != nil {
		return nil, err
	}
	return templatetypes, nil
}
