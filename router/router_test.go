package router_test

// import (
// 	"context"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/amren1254/stripe_integration/constant"
// 	router_mock "github.com/amren1254/stripe_integration/mocks/router"
// 	"github.com/amren1254/stripe_integration/router"
// 	"github.com/golang/mock/gomock"
// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// )

// func TestApp_InitRouteAndRun(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// Create a mock StripeHandler
// 	mockStripeHandler := router_mock.NewMockIRouter(ctrl) // NewMockStripeHandler(ctrl)

// 	// Set up the App instance with mock objects
// 	app := &router.App{
// 		Router:        mux.NewRouter(),
// 		StripeHandler: mockStripeHandler,
// 	}
// 	app.InitRoute(context.Background())

// 	// Example: Define test cases for expected routes and their handlers
// 	tests := []struct {
// 		name         string
// 		method       string
// 		path         string
// 		handlerFunc  func(http.ResponseWriter, *http.Request)
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Test Ping",
// 			method:       http.MethodGet,
// 			path:         constant.VERSION_ONE + constant.PING,
// 			handlerFunc:  mockStripeHandler.Ping,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Test CreateCheckoutSession",
// 			method:       http.MethodPost,
// 			path:         constant.VERSION_ONE + constant.CREATE_CHECKOUT_SESSION,
// 			handlerFunc:  mockStripeHandler.CreateCheckoutSession,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Test CreatePortalSession",
// 			method:       http.MethodPost,
// 			path:         constant.VERSION_ONE + constant.CREATE_PORTAL_SESSION,
// 			handlerFunc:  mockStripeHandler.CreatePortalSession,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Test WebHook",
// 			method:       http.MethodPost,
// 			path:         constant.VERSION_ONE + constant.WEBHOOK,
// 			handlerFunc:  mockStripeHandler.WebHook,
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create a request to the test server
// 			req, err := http.NewRequest(tt.method, tt.path, nil)
// 			if err != nil {
// 				t.Fatalf("could not create request: %v", err)
// 			}

// 			// Create a response recorder to record the response
// 			rr := httptest.NewRecorder()

// 			// Serve the HTTP request using the test server
// 			app.Router.ServeHTTP(rr, req)

// 			// Assert the response status code
// 			assert.Equal(t, tt.expectedCode, rr.Code, "unexpected status code")
// 		})
// 	}
// }
