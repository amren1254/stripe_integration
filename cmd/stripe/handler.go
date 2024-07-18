package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

type IStripHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	CreateCheckoutSession(w http.ResponseWriter, r *http.Request)
	CreatePortalSession(w http.ResponseWriter, r *http.Request)
	WebHook(w http.ResponseWriter, r *http.Request)
}

func (app *application) Ping(ctx *gin.Context) {
	// Set the JSON content type for the response
	ctx.Header("Content-Type", "application/json")

	// Write the HTTP status code and JSON response body
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (app *application) CreateCheckoutSession(ctx *gin.Context) {
	priceID := os.Getenv("PRODUCT_PRICE_ID") // get from config/env
	customerEmail := ctx.PostForm("customer_email")
	stripe.Key = app.config.stripe.key
	params := &stripe.CheckoutSessionParams{
		CustomerEmail: stripe.String(customerEmail),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String("http://localhost:9001/success.html"),
		ConsentCollection: &stripe.CheckoutSessionConsentCollectionParams{
			TermsOfService: stripe.String(string(stripe.CheckoutSessionConsentCollectionTermsOfServiceRequired)),
		},
		CustomText: &stripe.CheckoutSessionCustomTextParams{
			TermsOfServiceAcceptance: &stripe.CheckoutSessionCustomTextTermsOfServiceAcceptanceParams{
				Message: stripe.String("I agree to the [Terms of Service](https://example.com/terms)"),
			},
		},
	}
	session, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		ctx.String(http.StatusInternalServerError, "Failed to create checkout session")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"session_id":  session.ID,
		"session_url": session.URL,
	})
}

func (app *application) GetPaymentStatusHandler(ctx *gin.Context) {
	//get customer details from form value
	stripe.Key = app.config.stripe.key
	sessionId := ""
	params := stripe.CheckoutSessionParams{}
	result, err := session.Get(
		sessionId,
		&params,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, result)
	}
	ctx.JSON(http.StatusOK, result.PaymentStatus)
}
