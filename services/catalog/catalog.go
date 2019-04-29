package catalogs

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

// Create creates a new catalog and return the id
func (c *Client) Create(catalog *Catalog) (*Response, error) {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(catalog)

	if err != nil {
		return nil, err
	}

	body, err := c.HTTPRequest(fmt.Sprint("/v2/catalogs"), "POST", buf)

	response := &Response{}
	err = json.NewDecoder(body).Decode(catalog)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get gets the catalog by the Id
func (c *Client) Get(id string) (*Catalog, error) {

	body, err := c.HTTPRequest(fmt.Sprintf("/v2/catalogs/%v", id), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	catalog := &Catalog{}
	err = json.NewDecoder(body).Decode(catalog)

	if err != nil {
		return nil, err
	}
	return catalog, nil
}

// Update updates the values of an catalog
func (c *Client) Update(id string, catalog *Catalog) error {
	buf := bytes.Buffer{}
e\z
	err := json.NewEncoder(&buf).Encode(catalog)

	if err != nil {
		return err
	}
	_, err = c.HTTPRequest(fmt.Sprintf("/v2/catalogs/%v", id), "PUT", buf)

	if err != nil {
		return err
	}

	return nil
}

// Delete removes a catalog
func (c *Client) Delete(id string) error {
	_, err := c.HTTPRequest(fmt.Sprintf("/v2/catalogs/%v", id), "DELETE", bytes.Buffer{})

	if err != nil {
		return err
	}

	return nil
}
