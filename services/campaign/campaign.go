package campaign

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

// Create creates a new campaign and return the id
func (c *Client) Create(campaign *Campaign) (*Response, error) {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(campaign)

	if err != nil {
		return nil, err
	}

	body, err := c.HttpRequest(fmt.Sprint("/v2/campaigns"), "POST", buf)

	response := &Response{}
	err = json.NewDecoder(body).Decode(campaign)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// Get gets the campaign by the Id
func (c *Client) Get(id string) (*Campaign, error) {

	body, err := c.HttpRequest(fmt.Sprintf("/v2/campaigns/%v", id), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	campaign := &Campaign{}
	err = json.NewDecoder(body).Decode(campaign)

	if err != nil {
		return nil, err
	}
	return campaign, nil
}

// Update updates the values of an campaign
func (c *Client) Update(id string, campaign *Campaign) error {
	buf := bytes.Buffer{}

	err := json.NewEncoder(&buf).Encode(campaign)

	if err != nil {
		return err
	}
	_, err = c.HttpRequest(fmt.Sprintf("/v2/campaigns/%v", id), "PUT", buf)

	if err != nil {
		return err
	}

	return nil
}

// Delete removes a campaign
func (c *Client) Delete(id string) error {
	_, err := c.HttpRequest(fmt.Sprintf("/v2/campaigns/%v", id), "DELETE", bytes.Buffer{})

	if err != nil {
		return err
	}

	return nil
}
