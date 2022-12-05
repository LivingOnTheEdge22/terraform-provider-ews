package ews

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var baseURL string
var baseURLRev2 string
var baseURLAPI string

var baseURLews string

var descriptions map[string]string

func init() {

	baseURL = "https://my.incapsula.com/api/prov/v1"
	baseURLRev2 = "https://my.imperva.com/api/prov/v2"
	baseURLAPI = "https://api.imperva.com"

	baseURLews = "https://ews-management.abp-monsters.com"

	descriptions = map[string]string{
		"api_id": "The API identifier for API operations. You can retrieve this\n" +
			"from the Incapsula management console. Can be set via INCAPSULA_API_ID " +
			"environment variable.",

		"api_key": "The API key for API operations. You can retrieve this\n" +
			"from the Incapsula management console. Can be set via INCAPSULA_API_KEY " +
			"environment variable.",

		"base_url": "The base URL for API operations. Used for provider development.",
	}
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		APIID:   d.Get("api_id").(string),
		APIKey:  d.Get("api_key").(string),
		BaseURL: d.Get("base_url").(string),
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
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("EWS_BASE_URL", baseURL),
				Description: descriptions["base_url"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"incapsula_account_data": dataSourceAccount(),
		},

		ResourcesMap: map[string]*schema.Resource{

			"ews_site":    resourceSite(),
			"ews_account": resourceAccount(),
			"resourceEws": resourceEws(),
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
