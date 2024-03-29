# Psychz Provider Examples

This directory contains an examples of using Psychz services with Terraform.
Each example has its own README file containing more details on what it does.

Psychz Provider examples are grouped into following directories:

* [order_details](order_details/) - For calling order details API
* [order_express](order_express/) - For calling order express API
* [services_details](services_details/) - For calling services details API


## Using examples

To run any example, clone the repository, **adjust variables**, initialize plugins
and run `terraform apply` within the example's own directory.

```sh
git clone https://github.com/psychz-networks/terraform-provider-psychz.git
cd terraform-provider-psychz/examples/order_detail
terraform init
terraform apply
```
