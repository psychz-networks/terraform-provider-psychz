package psychz

import (
	"context"
	"encoding/json"

	// "fmt"
	// "net/http"
	// "strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	psychz "github.com/psychz-networks/terraform-provider-psychz/psychz/clients/dedicated_servers"
)

func DataPsychzOrderPlans() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrderPlansRead,
		Schema: map[string]*schema.Schema{
			"plan_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Plan ID to filter plans.",
			},
			"category_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Category ID to filter plans. Use 2 for dedicated category.",
			},
			"location_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Location code to filter plans. This value can be obtained using the location_list API.",
			},
			"billing_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Billing type to filter plans. Allowed values: '1=standard', '2=per usage'.",
			},
			"availability": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Availability status to filter plans. Allowed values: '1=all', '2=available', '3=out of stock'.",
			},
			"option_detail": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "By default its showing plan with configuration_items which are useful in /v1/order_create for config_option, to show plan listing without configuration_items please enter value '1'.",
			},

			"plans": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOrderPlansRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*psychz.Client)

	params := make(map[string]string)

	if planID, ok := d.GetOk("plan_id"); ok {
		params["plan_id"] = planID.(string)
	}

	if categoryID, ok := d.GetOk("category_id"); ok {
		params["category_id"] = categoryID.(string)
	}

	if locationCode, ok := d.GetOk("location_code"); ok {
		params["location_code"] = locationCode.(string)
	}

	if billingType, ok := d.GetOk("billing_type"); ok {
		params["billing_type"] = billingType.(string)
	}

	if availability, ok := d.GetOk("availability"); ok {
		params["availability"] = availability.(string)
	}

	if optionDetail, ok := d.GetOk("option_detail"); ok {
		params["option_detail"] = optionDetail.(string)
	}

	res, err := client.GetOrderPlans(ctx, params)
	if err != nil {
		return diag.FromErr(err)
	}

	// Marshal res into JSON format
	jsonData, err := json.Marshal(res)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("psychz_order_plans")
	d.Set("plans", string(jsonData))

	return nil
}
