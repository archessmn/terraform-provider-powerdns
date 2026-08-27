package main

import (
	"github.com/archessmn/terraform-provider-powerdns/powerdns"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: powerdns.Provider})
}
