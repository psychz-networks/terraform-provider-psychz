---
Subcategory: "Order Express"
---

## Resources


## Example Usage

``` 
resource "psychz_order_express" "express_server" {
  plan_id = plan id 
  order_quantity         = order quantity
  payment_mode           = payment mode
  auth_method            = auth method
  os_cat                 = os_category
  os_id                  = os id
  disk_partition_id      = disk partition id
  software_raid          = software raid
  hostname               =host name
  password               = password
  private_key            = private key
  partner_id             = partner id
  enforce_password_change = enforce password change
}
```

| Name | Type |
|------|------|
| psychz_order_express.express_server | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_disk_partition_id"></a> [disk\_partition\_id](#input\_disk\_partition\_id) | disk\_partition\_id can be obtained from api /os\_install\_category api | `number` | n/a | yes |
| <a name="input_enforce_password_change"></a> [enforce\_password\_change](#input\_enforce\_password\_change) | 0 - no password change, 1 - enforce password change option to require password change after first successful login. | `number` | n/a | yes |
| <a name="input_hostname"></a> [hostname](#input\_hostname) | hostname for os install (This field is required depends upon os templates, can check from api /os\_install\_category) | `string` | n/a | yes |
| <a name="input_order_id"></a> [order\_id](#input\_order\_id) | n/a | `any` | n/a | yes |
| <a name="input_order_quantity"></a> [order\_quantity](#input\_order\_quantity) | Order quantity. | `number` | n/a | yes |
| <a name="input_os_cat"></a> [os\_cat](#input\_os\_cat) | Os category 1 - Linux 2 - Windows | `number` | n/a | yes |
| <a name="input_os_id"></a> [os\_id](#input\_os\_id) | os\_id can be obtained from api /os\_install\_category api | `number` | n/a | yes |
| <a name="input_partner_id"></a> [partner\_id](#input\_partner\_id) | Partner ID (System will execute specific rules that apply to partner ID) | `number` | n/a | yes |
| <a name="input_password"></a> [password](#input\_password) | Password must contain 8-12 alphanumeric characters only. Special characters are not allowed (This field is required depends upon os templates, can check from api /os\_install\_category) | `string` | n/a | yes |
| <a name="input_payment_mode"></a> [payment\_mode](#input\_payment\_mode) | Payment mode 1 = 'credit apply' 2 = 'credit card' (*A default credit card is required on file). | `number` | n/a | yes |
| <a name="input_auth_method"></a> [auth\_method](#input\_auth\_method) | 1 - password, 2 - Private Key. | `number` | n/a | yes |
| <a name="input_plan_id"></a> [plan\_id](#input\_plan\_id) | Plan ID for the order. | `number` | n/a | yes |
| <a name="input_private_key"></a> [private\_key](#input\_private\_key) | private key (ssh-rsa public key) for os install if selected auth method 2  | `string` | n/a | yes |
| <a name="input_software_raid"></a> [software\_raid](#input\_software\_raid) | software\_raid can be obtained from api /os\_install\_category api (This field is required depends upon os templates, can check from api /os\_install\_category) | `number` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_server_order_express_status"></a> [server\_order\_express\_status](#output\_server\_order\_express\_status) | n/a |
| <a name="output_server_order_status"></a> [server\_order\_status](#output\_server\_order\_status) | n/a |
