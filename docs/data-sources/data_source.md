---
layout: "scaffolding"
page_title: "Scaffolding: scaffolding_data_source"
sidebar_current: "docs-scaffolding-data-source"
description: |-
  Sample data source in the Terraform provider scaffolding.
---

# scaffolding_data_source

Sample data source in the Terraform provider scaffolding.

## Example Usage

```hcl
data "scaffolding_data_source" "example" {
  sample_attribute = "foo"
}
```

## Attributes Reference

* `sample_attribute` - Sample attribute.
