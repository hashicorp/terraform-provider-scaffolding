package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceScaffolding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceScaffoldingRead,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceScaffoldingRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
