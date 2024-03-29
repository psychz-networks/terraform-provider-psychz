package psychz

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	psychz "github.com/psychz-networks/terraform-provider-psychz/psychz/clients/dedicated_servers"
)

func DataPsychzOrderDetail() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrderDetailRead,
		Schema: map[string]*schema.Schema{
			"order_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order id for fetching order details",
			},
			"order_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOrderDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*psychz.Client)
	orderID := d.Get("order_id").(int)

	orderDetail, err := client.GetOrderDetail(ctx, orderID)
	if err != nil {
		return diag.FromErr(err)
	}
	// Marshal res into JSON format
	jsonData, err := json.Marshal(orderDetail)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("psychz_order_detail")
	d.Set("order_info", string(jsonData))

	return nil
}
