---
page_title: "scaffolding_resource Resource - terraform-provider-scaffolding"
subcategory: ""
description: |-
  Sample resource in the Terraform provider scaffolding.
---

# Resource `scaffolding_resource`

Sample resource in the Terraform provider scaffolding.

## Example Usage

```terraform
resource "scaffolding_resource" "example" {
  sample_attribute = "foo"
}
```

## Schema

### Optional

- **id** (String, Optional) The ID of this resource.
- **sample_attribute** (String, Optional) Sample attribute.


