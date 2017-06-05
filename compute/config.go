package compute

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/errwrap"
)

type ConfigClient struct {
	*Compute
}

// Config represents configuration for your account.
type Config struct {
	// DefaultNetwork is the network that docker containers are provisioned on.
	DefaultNetwork string `json:"default_network"`
}

type GetConfigInput struct{}

// GetConfig outputs configuration for your account.
func (c *ConfigClient) GetConfig(ctx context.Context, input *GetConfigInput) (*Config, error) {
	path := fmt.Sprintf("/%s/config", c.client.AccountName)
	respReader, err := c.executeRequest(ctx, http.MethodGet, path, nil)
	if respReader != nil {
		defer respReader.Close()
	}
	if err != nil {
		return nil, errwrap.Wrapf("Error executing GetConfig request: {{err}}", err)
	}

	var result *Config
	decoder := json.NewDecoder(respReader)
	if err = decoder.Decode(&result); err != nil {
		return nil, errwrap.Wrapf("Error decoding GetConfig response: {{err}}", err)
	}

	return result, nil
}

type UpdateConfigInput struct {
	// DefaultNetwork is the network that docker containers are provisioned on.
	DefaultNetwork string `json:"default_network"`
}

// UpdateConfig updates configuration values for your account.
func (c *ConfigClient) UpdateConfig(ctx context.Context, input *UpdateConfigInput) (*Config, error) {
	path := fmt.Sprintf("/%s/config", c.client.AccountName)
	respReader, err := c.executeRequest(ctx, http.MethodPut, path, input)
	if respReader != nil {
		defer respReader.Close()
	}
	if err != nil {
		return nil, errwrap.Wrapf("Error executing UpdateConfig request: {{err}}", err)
	}

	var result *Config
	decoder := json.NewDecoder(respReader)
	if err = decoder.Decode(&result); err != nil {
		return nil, errwrap.Wrapf("Error decoding UpdateConfig response: {{err}}", err)
	}

	return result, nil
}
