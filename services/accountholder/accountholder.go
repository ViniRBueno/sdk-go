package accountholder

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

// Get gets the AccountHolder by the campaignid
func (c *Client) Get(id string) (*AccountHolder, error) {

	body, err := c.HTTPRequest(fmt.Sprintf("/v1/campaigns/%v/originAccountHolder", id), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	accountholder := &AccountHolder{}
	err = json.NewDecoder(body).Decode(accountholder)

	if err != nil {
		return nil, err
	}
	return accountholder, nil
}
