package ews

import (
	"errors"
	"log"
	"strings"
)

// Config represents the configuration required for the Ews Client
type Config struct {
	// API Identifier
	APIID string

	// API Key
	APIKey string

	// Base URL EWS (no trailing slash)
	// This endpoint is unlikely to change in the near future
	baseURLEWS string
}

var missingAPIIDMessage = "API Identifier (api_id) must be provided"
var missingAPIKeyMessage = "API Key (api_key) must be provided"
var missingBaseURLEWSMessage = "Base URL EWS must be provided"

// Client configures and returns a fully initialized Ews Client
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
		return nil, errors.New(missingBaseURLEWSMessage)
	}

	// Create client
	client := NewClient(c)

	return client, nil
}
