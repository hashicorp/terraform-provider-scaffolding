package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}
