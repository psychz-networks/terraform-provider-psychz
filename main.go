// main.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/psychz-networks/terraform-provider-psychz/psychz"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: psychz.Provider,
	})
}
