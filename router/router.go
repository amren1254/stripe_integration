package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/amren1254/stripe_integration/constant"
	"github.com/amren1254/stripe_integration/handler"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router        *gin.Engine
	StripeHandler *handler.StripeHandler
}

type IRouter interface {
	InitRoute(ctx context.Context)
	Run(ctx context.Context) error
}

func (app *App) InitRoute(ctx context.Context) {
	app.Router = gin.Default()

	v1 := app.Router.Group(constant.VERSION_ONE)
	{
		v1.GET(constant.PING, app.StripeHandler.Ping)
		v1.POST(constant.CREATE_CHECKOUT_SESSION, app.StripeHandler.CreateCheckoutSession)
		v1.POST(constant.CREATE_PORTAL_SESSION, app.StripeHandler.CreatePortalSession)
		v1.POST(constant.WEBHOOK, app.StripeHandler.WebHook)
	}
}

func (app *App) Run(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", constant.HOST_ADDRESS, constant.PORT_NUMBER)
	log.Printf("Listening on %s", addr)

	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		return err
	}

	return nil
}
