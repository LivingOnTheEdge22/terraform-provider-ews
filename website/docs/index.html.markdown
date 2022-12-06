---
layout: "incapsula"
page_title: "Provider: Incapsula"
sidebar_current: "docs-incapsula-index"
description: |-
  The Incapsula provider is used to interact with resources supported by Imperva. The provider needs to be configured with the proper credentials before it can be used.
---

# Incapsula Provider

The EWS provider is used to interact with resources supported by Imperva Edge services. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Incapsula provider
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

* `api_id` - (Required) The Incapsula API id associated with the account. This can also be
  specified with the `INCAPSULA_API_ID` shell environment variable.
* `api_key` - (Required) The Incapsula API key. This can also be specified with the 
  `INCAPSULA_API_KEY` shell environment variable.
