terraform {
  required_providers {
    ews = {
      source = "terraform-providers/ews"
      version = "1.0.7"
    }
  }
}
 
variable "ews_api_id" {
  type        = number
  description = "API ID"
}
 
variable "ews_api_key" {
  type        = string
  description = "API KEY"
}

provider "ews" {
  api_id = var.ews_api_id
  api_key = var.ews_api_key
  base_url_ews = "https://ews-management.abp-monsters.com"
}

resource "ews_lambda" "first_lambda" {
	account_id       = -1
	lambda           = "leaked-director.zip"
	lambda_name      = "leaked-redirector"
	filter_path      = "/login"
	deployed         = false
}