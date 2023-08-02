package psychz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	clients "github.com/psychz-networks/terraform-provider-psychz/psychz/clients/dedicated_servers"
	data_sources "github.com/psychz-networks/terraform-provider-psychz/psychz/data_sources/dedicated_servers"
	resources "github.com/psychz-networks/terraform-provider-psychz/psychz/resources/dedicated_servers"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PSYCHZ_ACCESS_USERNAME", nil),
				Description: "Access username for the Psychz Networks API.",
			},
			"access_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("PSYCHZ_ACCESS_TOKEN", nil),
				Description: "Access token for the Psychz Networks API.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"psychz_order_express": resources.ResourceOrderExpress(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"psychz_order_detail":   data_sources.DataPsychzOrderDetail(),
			"psychz_order_plans":    data_sources.DataPsychzOrderPlans(),
			"psychz_service_detail": data_sources.DataPsychzServiceDetail(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	accessToken := d.Get("access_token").(string)
	accessUsername := d.Get("access_username").(string)

	var diags diag.Diagnostics

	if accessToken == "" || accessUsername == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Access token and username must be provided",
			Detail:   "Access token and username are required to access the Psychz API.",
		})
		return nil, diags
	}

	return clients.NewClient(accessToken, accessUsername), diags
}
