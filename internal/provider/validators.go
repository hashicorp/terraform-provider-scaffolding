package provider

import (
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateStringIs(expected string) schema.SchemaValidateDiagFunc {
	return func(v interface{}, p cty.Path) diag.Diagnostics {
		value := v.(string)
		var diags diag.Diagnostics
		if value != expected {
			diag := diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "wrong value",
				Detail:   fmt.Sprintf("%q is not %q", value, expected),
			}
			diags = append(diags, diag)
		}
		return diags
	}
}
