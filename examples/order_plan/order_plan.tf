variable "plan_id" {}
variable "category_id" {}
variable "location_code" {}
variable "billing_type" {}
variable "availability" {}
variable "option_detail" {}

data "psychz_order_plans" "server" {
  plan_id        = var.plan_id
  category_id    = var.category_id
  location_code  = var.location_code
  billing_type   = var.billing_type
  availability   = var.availability
  option_detail  = var.option_detail
}

output "server_order_plan" {
  value = data.psychz_order_plans.server.plans
}
