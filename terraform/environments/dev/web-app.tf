module "webservice_stripe_integration"{
    source = "../../modules/stripe-integration"
    resource_group_name="mosspark-rg"
    resource_group_location="East US"
    app_service_plan_name = "mosspark-app-service-sp"
    app_service_plan_sku_name="B1"
    app_service_name="stripe_integration"
}