package provider

import (
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func validateStringIsABC(v interface{}, p cty.Path) diag.Diagnostics {
	value := v.(string)
	var diags diag.Diagnostics
	if value != "abc" {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "wrong value",
			Detail:   fmt.Sprintf("%q is not abc", value),
		}
		diags = append(diags, diag)
	}
	return diags
}
