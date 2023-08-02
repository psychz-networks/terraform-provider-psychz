variable "order_id" {}
data "psychz_order_detail" "server" {
  order_id = var.order_id
}

output "server_order_status" {
  value = data.psychz_order_detail.server.order_info
}
