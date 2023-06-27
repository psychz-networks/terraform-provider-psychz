package psychz

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	psychz "github.com/psychz-networks/terraform-provider-psychz/psychz/clients/dedicated_servers"
)

func ResourceOrderExpress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrderExpressCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,
		UpdateContext: schema.NoopContext,
		Schema: map[string]*schema.Schema{
			"plan_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Plan ID for the order.",
			},
			"order_quantity": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Order quantity.",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value != 1 {
						errors = append(errors, fmt.Errorf("%q must be 1", k))
					}
					return
				},
			},
			"payment_mode": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Payment mode 1 = 'credit apply' 2 = 'credit card' (*A default credit card is required on file).",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value != 1 && value != 2 {
						errors = append(errors, fmt.Errorf("%q must be either 1 or 2", k))
					}
					return
				},
			},
			"os_cat": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "OS category 1 - Linux 2 - Windows",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value != 1 && value != 2 {
						errors = append(errors, fmt.Errorf("%q must be either 1 or 2", k))
					}
					return
				},
			},
			"os_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "OS ID obtained from /os_install_category API",
			},
			"disk_partition_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Disk partition ID obtained from /os_install_category API",
			},
			"software_raid": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Software RAID obtained from /os_install_category API",
			},
			"auth_method": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "auth method for OS install must be either 1 or 2",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value != 1 && value != 2 {
						errors = append(errors, fmt.Errorf("%q must be either 1 or 2", k))
					}
					return
				},
			},
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Hostname for OS install",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password for OS install",
			},
			"private_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Private key for OS install",
			},
			"partner_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Partner ID for OS install",
			},
			"enforce_password_change": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Enforce password change for OS install",
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(int)
					if value != 0 && value != 1 {
						errors = append(errors, fmt.Errorf("%q must be either 0 or 1", k))
					}
					return
				},
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The order message.",
			},
			"client_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The client ID.",
			},
		},
	}
}

func resourceOrderExpressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*psychz.Client)

	authMethod := d.Get("auth_method").(int)
	password := d.Get("password").(string)
	private_key := d.Get("private_key").(string)

	if authMethod == 1 && password == "" {
		return diag.Errorf("Password is required for this auth method")
	} else if authMethod == 2 && private_key == "" {
		return diag.Errorf("Private key is required for this auth method")
	}

	// Prepare the order details
	orderData := map[string]interface{}{
		"plan_id":                 d.Get("plan_id").(int),
		"order_quantity":          d.Get("order_quantity").(int),
		"payment_mode":            d.Get("payment_mode").(int),
		"os_cat":                  d.Get("os_cat").(int),
		"os_id":                   d.Get("os_id").(int),
		"disk_partition_id":       d.Get("disk_partition_id").(int),
		"auth_method":             d.Get("auth_method").(int),
		"software_raid":           d.Get("software_raid").(int),
		"hostname":                d.Get("hostname").(string),
		"password":                d.Get("password").(string),
		"private_key":             d.Get("private_key").(string),
		"partner_id":              d.Get("partner_id").(int),
		"enforce_password_change": d.Get("enforce_password_change").(int),
	}

	// Make the API call
	resp, err := client.CreateOrderExpress(ctx, orderData)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the resource ID and order_id
	d.SetId(resp.Data.OrderID)
	d.Set("message", resp.Data.Message)
	d.Set("client_id", resp.Data.ClientID)

	return nil
}
