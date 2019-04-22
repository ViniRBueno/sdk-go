package tenantdiscovery

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TenantClient is the client for Tenant Discovery Rest Api
type TenantClient struct {
	BaseClient
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, token string) *TenantClient {

	return &TenantClient{
		NewBase(
			hostname,
			token,
		)}
}

// Get gets an tenant configuration for campaign
func (c *TenantClient) Get(id string) (*Tenant, error) {

	body, err := c.httpRequest(fmt.Sprintf("/%v/.well-known/tenant-configuration", id), "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	tenant := &Tenant{}
	err = json.NewDecoder(body).Decode(tenant)

	if err != nil {
		return nil, err
	}

	return tenant, nil
}

// Create creates a new tenant configuraton for campaign
func (c *TenantClient) Create(tenant *Tenant) error {

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(tenant)

	if err != nil {
		return err
	}

	_, err = c.httpRequest(fmt.Sprintf("/tenant-configuration/%v", tenant.ID), "POST", buf)

	if err != nil {
		return err
	}

	return nil
}

// Update updates the values of an tenant discovery
func (c *TenantClient) Update(tenant *Tenant) error {
	buf := bytes.Buffer{}

	err := json.NewEncoder(&buf).Encode(tenant.Auth)

	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("/tenant-configuration/%v/auth", tenant.ID), "PUT", buf)

	if err != nil {
		return err
	}
	return nil
}

// Delete removes an tenant discovery configuration
func (c *TenantClient) Delete(id string) error {

	_, err := c.httpRequest(fmt.Sprintf("/tenant-configuration/%v", id), "DELETE", bytes.Buffer{})

	if err != nil {
		return err
	}

	return nil
}
