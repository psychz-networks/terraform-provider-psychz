package psychz

import (
	"context"
	"strconv"

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
			"status": {
				Type:        schema.TypeString,
				Description: "The API Call status.",
				Computed:    true,
			},
			"order_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The order time.",
			},
			"last_activity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The order last activity.",
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

	d.SetId(strconv.Itoa(orderID))
	d.Set("status", orderDetail.Data.OrderStatus)
	d.Set("order_time", orderDetail.Data.OrderTime)
	d.Set("last_activity", orderDetail.Data.LastActivity)

	return nil
}
