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

// Create creates a new template and return the id
func (c *Client) Create(template *Template) (*Response, error) {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(template)

	if err != nil {
		return nil, err
	}

	body, err := c.HTTPRequest(fmt.Sprint("/v2/template"), "POST", buf)

	response := &Response{}
	err = json.NewDecoder(body).Decode(template)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get gets the template by the Id
func (c *Client) Get(id string) (*Template, error) {

	body, err := c.HTTPRequest(fmt.Sprintf("/v2/template/%v", id), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	template := &Template{}
	err = json.NewDecoder(body).Decode(template)

	if err != nil {
		return nil, err
	}
	return template, nil
}

// Update updates the values of an template
func (c *Client) Update(id string, template *Template) error {
	buf := bytes.Buffer{}

	err := json.NewEncoder(&buf).Encode(template)

	if err != nil {
		return err
	}
	_, err = c.HTTPRequest(fmt.Sprintf("/v2/template/%v", id), "PUT", buf)

	if err != nil {
		return err
	}

	return nil
}

// Delete removes a template
func (c *Client) Delete(id string) error {
	_, err := c.HTTPRequest(fmt.Sprintf("/v2/template/%v", id), "DELETE", bytes.Buffer{})

	if err != nil {
		return err
	}

	return nil
}
