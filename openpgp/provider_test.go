package openpgp

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}

var testProviders = map[string]*schema.Provider{
	"openpgp": Provider(),
}
