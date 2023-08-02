variable "service_id" {}
data "psychz_service_detail" "server" {
  service_id = var.service_id
}

output "server_service_detail" {
  value = data.psychz_service_detail.server
}
