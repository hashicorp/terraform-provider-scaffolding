package provider

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

func TestValidateStringIsABC(t *testing.T) {
	v := "abc"
	if diags := validateStringIs("abc")(v, *new(cty.Path)); len(diags) != 0 {
		t.Fatalf("%q should be abc", v)
	}

	v = "not-abc"
	if diags := validateStringIs("abc")(v, *new(cty.Path)); len(diags) == 0 {
		t.Fatalf("%q should NOT be abc", v)
	}
}
