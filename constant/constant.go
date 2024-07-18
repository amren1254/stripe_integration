package constant

import "os"

// endpoint constants
const (
	ENV_PATH                = "../../.env"
	PING                    = "/ping"
	FORWARD_SLASH           = "/"
	VERSION_ONE             = "/v1"
	CREATE_CHECKOUT_SESSION = "/create-checkout-session"
	GET_PAYMENT_STATUS      = "/get-payment-status"
)

var (
	PRODUCT_PRICE_ID = os.Getenv("PRODUCT_PRICE_ID")
)
