package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	portalsession "github.com/stripe/stripe-go/v79/billingportal/session"
	"github.com/stripe/stripe-go/v79/checkout/session"
	"github.com/stripe/stripe-go/webhook"
)

type IStripHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	CreateCheckoutSession(w http.ResponseWriter, r *http.Request)
	CreatePortalSession(w http.ResponseWriter, r *http.Request)
	WebHook(w http.ResponseWriter, r *http.Request)
}

type StripeHandler struct {
}

func (sh *StripeHandler) Ping(ctx *gin.Context) {
	// Set the JSON content type for the response
	ctx.Header("Content-Type", "application/json")

	// Write the HTTP status code and JSON response body
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (sh *StripeHandler) CreateCheckoutSession(ctx *gin.Context) {
	priceID := os.Getenv("PRODUCT_PRICE_ID") // get from config/env
	stripe.Key = os.Getenv("STRIPE_KEY")
	customerEmail := ctx.PostForm("customer_email")
	params := &stripe.CheckoutSessionParams{
		CustomerEmail: stripe.String(customerEmail),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSubscription)),
		SuccessURL: stripe.String("http://localhost:8080/success.html"),
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
		"session_url": session.URL,
	})
}

func (sh *StripeHandler) CreatePortalSession(ctx *gin.Context) {
	stripe.Key = os.Getenv("STRIPE_KEY")
	err := ctx.Request.ParseForm()
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to parse form")
		return
	}

	sessionID := ctx.PostForm("session_id")
	fmt.Println(sessionID)

	s, err := session.Get(sessionID, nil)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to retrieve session")
		log.Printf("session.Get: %v", err)
		return
	}

	params := &stripe.BillingPortalSessionParams{
		Customer: stripe.String(s.Customer.ID),
	}

	ps, err := portalsession.New(params)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to create portal session")
		log.Printf("billingportal.New: %v", err)
		return
	}
	ctx.JSON(http.StatusCreated, ps.URL)
}

func (sh *StripeHandler) WebHook(ctx *gin.Context) {
	const MaxBodyBytes = int64(65536)
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MaxBodyBytes)
	payload, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		ctx.String(http.StatusServiceUnavailable, "Failed to read request body")
		return
	}

	// Replace this endpoint secret with your endpoint's unique secret
	// If you are testing with the CLI, find the secret by running 'stripe listen'
	// If you are using an endpoint defined with the API or dashboard, look in your webhook settings
	// at https://dashboard.stripe.com/webhooks
	endpointSecret := "whsec_12345"
	signatureHeader := ctx.GetHeader("Stripe-Signature")
	event, err := webhook.ConstructEvent(payload, signatureHeader, endpointSecret)
	if err != nil {
		log.Printf("⚠️  Webhook signature verification failed. %v\n", err)
		ctx.String(http.StatusBadRequest, "Webhook signature verification failed")
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "customer.subscription.deleted":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v\n", err)
			ctx.String(http.StatusBadRequest, "Error parsing webhook JSON")
			return
		}
		log.Printf("Subscription deleted for %s.", subscription.ID)
		// Then define and call a func to handle the deleted subscription.
		// handleSubscriptionCanceled(subscription)
	case "customer.subscription.updated":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v\n", err)
			ctx.String(http.StatusBadRequest, "Error parsing webhook JSON")
			return
		}
		log.Printf("Subscription updated for %s.", subscription.ID)
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handleSubscriptionUpdated(subscription)
	case "customer.subscription.created":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v\n", err)
			ctx.String(http.StatusBadRequest, "Error parsing webhook JSON")
			return
		}
		log.Printf("Subscription created for %s.", subscription.ID)
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handleSubscriptionCreated(subscription)
	case "customer.subscription.trial_will_end":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v\n", err)
			ctx.String(http.StatusBadRequest, "Error parsing webhook JSON")
			return
		}
		log.Printf("Subscription trial will end for %s.", subscription.ID)
		// Then define and call a func to handle the successful attachment of a PaymentMethod.
		// handleSubscriptionTrialWillEnd(subscription)
	case "entitlements.active_entitlement_summary.updated":
		var subscription stripe.Subscription
		err := json.Unmarshal(event.Data.Raw, &subscription)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v\n", err)
			ctx.String(http.StatusBadRequest, "Error parsing webhook JSON")
			return
		}
		log.Printf("Active entitlement summary updated for %s.", subscription.ID)
		// Then define and call a func to handle active entitlement summary updated.
		// handleEntitlementUpdated(subscription)
	default:
		log.Printf("Unhandled event type: %s\n", event.Type)
		ctx.String(http.StatusBadRequest, "Unhandled event type")
		return
	}

	// Respond with 200 OK status
	ctx.String(http.StatusOK, "Webhook received")
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
