---
Subcategory: "Order Detail"
---

## Resources

## Example Usage

``` 
variable "order_id" {}
data "psychz_order_detail" "server" {
  order_id = var.order_id
}
```

| Name | Type |
|------|------|
| psychz_order_detail.server | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="order_id"></a> [order\_id](#input\_order\_id) | Order ID (System will execute specific rules that apply to order ID) | `number` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_server_order_status"></a> [server\_order\_status](#output\_server\_order\_status) | n/a |
