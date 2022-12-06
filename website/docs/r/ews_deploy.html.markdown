---
layout: "ews"
page_title: "EWS: ews_deploy"
sidebar_current: "docs-ews-resource-ews_deploy"
description: |- Provides an EWS Deploy lambda resource.
---

# ews_compile

Deploy a previously compiled lambda. 
Note that import and delete actions are temporary unavailable.

## Example Usage

### Basic Usage - Ews Compile

```hcl
resource "ews_lambda_deploy" "director_lambda_deploy" {
  account_id       = -1
  lambda_name      = ews_lambda_compile.director_lambda.lambda_name
  filter_path      = "/login"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) Account to operate on .
* `lambda_name` - (Optional) Lambda name
* `filter_path` - (Optional) A site path to deploy the lambda on.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the Site Monitoring. The id is identical to Site id.
