---
layout: "ews"
subcategory: "Compile"
page_title: "EWS: ews_lambda_compile"
sidebar_current: "docs-ews-resource-ews_lambda_compile"
description: |-
  Provides an EWS Compile lambda resource.
---

# ews_lambda_compile

Upload and compile lambda. Note that import and delete actions are temporary unavailable.

## Example Usage

### Basic Usage - Ews Lambda Compile

```hcl
resource "ews_lambda_compile" "director_lambda_compile" {
  account_id       = -1
  site_id          = 444444
  lambda_name      = "leaked-redirector"
  lambda_path      = "../leaked-redirector/leaked-redirector.zip"
}
```

## Argument Reference

The following arguments are supported:

* `account_id`  - (Required) Account to operate on.
* `site_id`     - (Required) Numeric identifier of the site.
* `lambda_name` - (Required) Lambda name.
* `lambda_path` - (Required) Path to a local zip containing all files needed to compile lambda.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the Site Monitoring. The id is identical to Site id.
