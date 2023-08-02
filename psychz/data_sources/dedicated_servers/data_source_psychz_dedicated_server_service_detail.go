package psychz

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	psychz "github.com/psychz-networks/terraform-provider-psychz/psychz/clients/dedicated_servers"
)

func DataPsychzServiceDetail() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServiceDetailRead,
		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order id for fetching order details",
			},
			"ip_assignments": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServiceDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*psychz.Client)
	serviceID := d.Get("service_id").(int)

	ServiceDetail, err := client.GetServiceDetail(ctx, serviceID)
	if err != nil {
		return diag.FromErr(err)
	}
	// Marshal res into JSON format
	jsonData, err := json.Marshal(ServiceDetail)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("psychz_service_detail")
	d.Set("ip_assignments", string(jsonData))

	return nil
}
