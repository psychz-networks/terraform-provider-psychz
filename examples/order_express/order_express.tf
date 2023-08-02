variable "plan_id" {
  description = "Plan ID for the order."
  type        = number
}

variable "order_quantity" {
  description = "Order quantity."
  type        = number
}

variable "payment_mode" {
  description = "Payment mode 1 = 'credit apply' 2 = 'credit card' (*A default credit card is required on file)."
  type        = number
}

variable "os_cat" {
  description = "Os category 1 - Linux 2 - Windows"
  type        = number
}

variable "os_id" {
  description = "os_id can be obtained from api /os_install_category api"
  type        = number
}

variable "disk_partition_id" {
  description = "disk_partition_id can be obtained from api /os_install_category api"
  type        = number
}

variable "software_raid" {
  description = "software_raid can be obtained from api /os_install_category api (This field is required depends upon os templates, can check from api /os_install_category)"
  type        = number
}
variable "auth_method" {
  description = "	1 - Password, 2 - Private Key."
  type        = number
}

variable "hostname" {
  description = "Hostname for os install (This field is required depends upon os templates, can check from api /os_install_category)"
  type        = string
}

variable "password" {
  description = "Password must contain 12-30 alphanumeric characters only. Special characters are  allowed (This field is required depends upon os templates, can check from api /os_install_category)"
  type        = string
}
variable "private_key" {
  description = "private_key for os install"
  type        = string
}
variable "partner_id" {
  description = "Partner ID (System will execute specific rules that apply to partner ID)"
  type        = number
}

variable "enforce_password_change" {
  description = "	0 - no password change, 1 - enforce password change option to require password change after first successful login."
  type        = number
}

resource "psychz_order_express" "express_server" {
  plan_id = var.plan_id
  order_quantity         = var.order_quantity
  payment_mode           = var.payment_mode
  auth_method            = var.auth_method
  os_cat            = var.os_cat
  os_id                  = var.os_id
  disk_partition_id      = var.disk_partition_id
  software_raid          = var.software_raid
  hostname               = var.hostname
  password               = var.password
  private_key            = var.private_key
  partner_id             = var.partner_id
  enforce_password_change = var.enforce_password_change
}

output "server_order_express_status" {
  value = psychz_order_express.express_server
}


