output "app_service_url" {
  value = "https://${azurerm_linux_web_app.stripe_integration_app_service.default_hostname}"
}