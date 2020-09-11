package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaffolding() *schema.Resource {
	return &schema.Resource{
		Create: resourceScaffoldingCreate,
		Read:   resourceScaffoldingRead,
		Update: resourceScaffoldingUpdate,
		Delete: resourceScaffoldingDelete,

		Schema: map[string]*schema.Schema{
			"sample_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceScaffoldingCreate(d *schema.ResourceData, meta interface{}) error {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}

func resourceScaffoldingRead(d *schema.ResourceData, meta interface{}) error {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}

func resourceScaffoldingUpdate(d *schema.ResourceData, meta interface{}) error {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}

func resourceScaffoldingDelete(d *schema.ResourceData, meta interface{}) error {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}
