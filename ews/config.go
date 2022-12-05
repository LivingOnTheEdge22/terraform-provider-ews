package ews

import (
	"errors"
	"log"
	"strings"
)

// Config represents the configuration required for the Incapsula Client
type Config struct {
	// API Identifier
	APIID string

	// API Key
	APIKey string

	// Base URL EWS (no trailing slash)
	// This endpoint is unlikely to change in the near future
	baseURLEWS string

	// Base URL (no trailing slash)
	// This endpoint is unlikely to change in the near future
	BaseURL string

	// Base URL Revision 2 (no trailing slash)
	// Updates to APIv1 are underway and newer resources are supported
	// Rev2 includes the move to Swagger, appropriate method verbs (not everything is a post)
	// The other endpoints will eventually move over but we'll need the following for now
	BaseURLRev2 string

	// Base URL API
	// API V2
	// Same as revision 2 but with a different subdomain
	BaseURLAPI string
}

var missingAPIIDMessage = "API Identifier (api_id) must be provided"
var missingAPIKeyMessage = "API Key (api_key) must be provided"
var missingBaseURLMessage = "Base URL EWS must be provided"

// Client configures and returns a fully initialized Incapsula Client
func (c *Config) Client() (interface{}, error) {
	log.Println("[INFO] Checking API credentials for client instantiation")

	// Check API Identifier
	if strings.TrimSpace(c.APIID) == "" {
		return nil, errors.New(missingAPIIDMessage)
	}

	// Check API Key
	if strings.TrimSpace(c.APIKey) == "" {
		return nil, errors.New(missingAPIKeyMessage)
	}

	// Check Base URL
	if strings.TrimSpace(c.baseURLEWS) == "" {
		return nil, errors.New(missingBaseURLMessage)
	}

	// Create client
	client := NewClient(c)

	return client, nil
}
