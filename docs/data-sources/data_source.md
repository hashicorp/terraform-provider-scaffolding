---
page_title: "scaffolding_data_source Data Source - terraform-provider-scaffolding"
subcategory: ""
description: |-
  Sample data source in the Terraform provider scaffolding.
---

# Data Source `scaffolding_data_source`

Sample data source in the Terraform provider scaffolding.

## Example Usage

```terraform
data "scaffolding_data_source" "example" {
  sample_attribute = "foo"
}
```

## Schema

### Required

- **sample_attribute** (String, Required) Sample attribute.

### Optional

- **id** (String, Optional) The ID of this resource.


