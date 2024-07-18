package main

import (
	"github.com/amren1254/stripe_integration/constant"
	"github.com/gin-gonic/gin"
)

// type IRouter interface {
// 	Route(ctx context.Context)
// }

func (app *application) route() *gin.Engine {
	routes := gin.Default()
	if app.config.env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	v1 := routes.Group(constant.VERSION_ONE)
	{
		v1.GET(constant.PING, app.Ping)
		v1.POST(constant.CREATE_CHECKOUT_SESSION, app.CreateCheckoutSession)
		v1.POST(constant.CREATE_PORTAL_SESSION, app.CreatePortalSession)
		v1.POST(constant.WEBHOOK, app.WebHook)
	}
	return routes
}
