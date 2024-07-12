variable "resource_group_name" {
  description = "Resource Group name"
  type        = string
}
variable "resource_group_location" {
  description = "Resource Group Location"
  type        = string
}

variable "app_service_plan_name" {
  description = "Plan name for app service for stripe integration"
  type        = string
}

variable "app_service_plan_sku_name" {
  description = "sku size for app"
  type        = string
}

variable "app_service_name" {
  description = "Name for web app service to be deployed as azure web app"
  type        = string
}

variable "artifact_path" {
  description = "zip file path for web service"
  type        = string
}
