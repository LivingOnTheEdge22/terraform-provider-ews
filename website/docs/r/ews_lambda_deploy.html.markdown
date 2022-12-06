---
layout: "ews"
page_title: "EWS: ews_lambda_deploy"
sidebar_current: "docs-ews-resource-ews_lambda_deploy"
description: |- Provides an EWS Deploy lambda resource.
---

# ews_lambda_deploy

Deploy a previously compiled lambda. 
Note that import and delete actions are temporary unavailable.
Use `depends_on` to ensure compile will run before deploy.

## Example Usage

### Basic Usage - Ews Lambda Deploy

```hcl
resource "ews_lambda_deploy" "director_lambda_deploy" {
  account_id       = -1
  site_id          = 444444
  lambda_name      = "leaked-redirector"
  filter_path      = "/login"
  depends_on       = [ews_lambda_compile.director_lambda_compile]
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Required) Account to operate on.
* `site_id` - (Required) Numeric identifier of the site.
* `lambda_name` - (Required) Lambda name.
* `filter_path` - (Required) A site path to deploy the lambda on.

## Attributes Reference

The following attributes are exported:

* `id` - Unique identifier in the API for the Site Monitoring. The id is identical to Site id.
