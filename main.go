package main

import (
    "github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/robjporter/terraform-provider-f5ltprulesearch/f5ltprulesearch"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return f5ltprulesearch.Provider()
		},
	})
}