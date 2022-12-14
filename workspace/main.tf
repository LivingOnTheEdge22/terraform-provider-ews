terraform {
  required_providers {
    ews = {
      source = "terraform-providers/ews"
      version = "1.2.0"
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

resource "ews_lambda_compile" "director_lambda_compile" {
  account_id       = -1
  site_id          = "4444444"
  lambda_name      = "leaked-redirector"
  lambda_path      = "../../../../leaked-redirector/leaked-redirector.zip"
}

resource "ews_lambda_deploy" "director_lambda_deploy" {
  account_id       = -1
  site_id          = "4444444"
  lambda_name      = "leaked-redirector"
  filter_path      = "/login"
  depends_on = [ews_lambda_compile.director_lambda_compile]
}