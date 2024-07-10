package main

import (
	"context"
	"log"

	"github.com/amren1254/stripe_integration/handler"
	"github.com/amren1254/stripe_integration/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		return
	}

	server := router.App{
		Router:        &gin.Engine{},
		StripeHandler: &handler.StripeHandler{},
	}
	server.InitRoute(ctx)
	err = server.Run(ctx)
	if err != nil {
		log.Println(err)
	}
}
