package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceScaffolding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceScaffoldingRead,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceScaffoldingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	return diag.Errorf("not implemented")
}
