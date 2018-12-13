package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jtopjian/terraform-provider-jerakia/jerakia"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: jerakia.Provider,
	})
}
