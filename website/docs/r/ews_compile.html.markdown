---
layout: "ews"
page_title: "EWS: ews_compile"
sidebar_current: "docs-ews-resource-ews_compile"
description: |- Provides an EWS Compile lambda resource.
---

# ews_lambda_compile

Upload and compile lambda. Note that import and delete actions are temporary unavailable.

## Example Usage

### Basic Usage - Ews Compile

```hcl
resource "ews_lambda_compile" "director_lambda_compile" {
  account_id       = -1
  lambda_name      = "leaked-redirector"
  lambda_path      = "./leaked-director.zip"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) Account to operate on .
* `lambda_name` - (Optional) Lambda name 
* `lambda_path` - (Optional) Path to a local zip containing all files needed to compile lambda.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the Site Monitoring. The id is identical to Site id.
