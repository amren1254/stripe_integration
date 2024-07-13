resource "azurerm_service_plan" "app_service_plan" {
  name                = var.app_service_plan_name
  location            = azurerm_resource_group.app_service_resource_group.location
  resource_group_name = azurerm_resource_group.app_service_resource_group.name
  os_type             = "Linux"
  sku_name            = var.app_service_plan_sku_name

  tags = {
    environment = "dev"
  }
}


resource "azurerm_linux_web_app" "stripe_integration_app_service" {
  name                = var.app_service_name
  location            = azurerm_resource_group.app_service_resource_group.location
  resource_group_name = azurerm_resource_group.app_service_resource_group.name
  service_plan_id     = azurerm_service_plan.app_service_plan.id
  https_only          = true

  site_config {
    minimum_tls_version = "1.2"
  }
  lifecycle {
    ignore_changes = [service_plan_id]
  }
  app_settings = {
    WEBSITE_RUN_FROM_PACKAGE = 1
  }
}
