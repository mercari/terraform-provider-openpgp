package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"

	"github.com/mercari/terraform-provider-openpgp/openpgp"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: openpgp.Provider,
	})
}
