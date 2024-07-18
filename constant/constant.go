package constant

import "os"

// endpoint constants
const (
	ENV_PATH                = "../../.env"
	PING                    = "/ping"
	FORWARD_SLASH           = "/"
	VERSION_ONE             = "/v1"
	CREATE_CHECKOUT_SESSION = "/create-checkout-session"
	CREATE_PORTAL_SESSION   = "/create-portal-session"
	WEBHOOK                 = "/webhook"
)

var (
	PRODUCT_PRICE_ID = os.Getenv("PRODUCT_PRICE_ID")
)
