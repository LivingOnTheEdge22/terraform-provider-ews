package ews

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var baseURLEWS string

var descriptions map[string]string

func init() {

	baseURLEWS = "https://ews-management.abp-monsters.com"

	descriptions = map[string]string{
		"api_id": "The API identifier for API operations. You can retrieve this\n" +
			"from the Ews management console. Can be set via EWS_API_ID " +
			"environment variable.",

		"api_key": "The API key for API operations. You can retrieve this\n" +
			"from the Ews management console. Can be set via EWS_API_KEY " +
			"environment variable.",

		"base_url_ews": "The base URL for API operations. Used for provider development.",
	}
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		APIID:      d.Get("api_id").(string),
		APIKey:     d.Get("api_key").(string),
		baseURLEWS: d.Get("base_url_ews").(string),
	}

	return config.Client()
}

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EWS_API_ID", ""),
				Description: descriptions["api_id"],
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EWS_API_KEY", ""),
				Description: descriptions["api_key"],
			},
			"base_url_ews": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EWS_BASE_URL", baseURLEWS),
				Description: descriptions["base_url_ews"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"ews_lambda_compile": resourceEwsCompile(),
			"ews_lambda_deploy":  resourceEwsDeploy(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}
