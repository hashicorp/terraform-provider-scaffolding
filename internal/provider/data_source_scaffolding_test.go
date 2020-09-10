package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceScaffolding(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceScaffolding,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"scaffolding_data_source.foo", "sample_attribute", regexp.MustCompile("^ba")),
				),
			},
		},
	})
}

const testAccDataSourceScaffolding = `
resource "scaffolding_data_source" "foo" {
  sample_attribute = "bar"
}
`
