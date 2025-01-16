package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/vmware/terraform-provider-vcfa/vcfa"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: vcfa.Provider})
}
