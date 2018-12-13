package jerakia

import (
	"net/http"

	"github.com/jerakia/go-jerakia"
)

// Config represents a configuration struct for Jerakia.
type Config struct {
	// apiUrl is the URL to the Jerakia API service.
	apiUrl string

	// apiToken is the token to authenticate to Jerakia.
	apiToken string

	// client is the jerakia client.
	client jerakia.Client
}

// LoadAndValidate is a method used to initiate a client.
func (c *Config) LoadAndValidate() error {
	jConfig := jerakia.ClientConfig{
		URL:   c.apiUrl,
		Token: c.apiToken,
	}

	c.client = jerakia.NewClient(http.DefaultClient, jConfig)

	return nil
}
