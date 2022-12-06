---
layout: "ews"
page_title: "Provider: Ews"
sidebar_current: "docs-ews-index"
description: |-
  The Ews provider is used to interact with resources supported by Imperva. The provider needs to be configured with the proper credentials before it can be used.
---

# Ews Provider

The EWS provider is used to interact with resources supported by Imperva Edge services. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Ews provider
provider "ews" {
  api_id = var.ews_api_id
  api_key = var.ews_api_key
  base_url_ews = "https://ews-management.abp-monsters.com"
}

# Upload Lambda
resource "ews_lambda_compile" "director_lambda_compile" {
  account_id       = -1
  lambda_name      = "leaked-redirector"
  lambda_path      = "./leaked-director.zip"
}

# Deploy Lambda
resource "ews_lambda_deploy" "director_lambda_deploy" {
  account_id       = -1
  lambda_name      = ews_lambda_compile.director_lambda.lambda_name
  filter_path      = "/login"
}

```

## Argument Reference

The following arguments are supported:

* `api_id` - (Required) The Ews API id associated with the account. This can also be
  specified with the `EWS_API_ID` shell environment variable.
* `api_key` - (Required) The Ews API key. This can also be specified with the 
  `EWS_API_KEY` shell environment variable.
