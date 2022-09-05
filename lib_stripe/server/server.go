// server.go
//
// Use this sample code to handle webhook events in your integration.
//
// 1) Create a new Go module
//   go mod init example.com/stripe/webhooks/example
//
// 2) Paste this code into a new file (server.go)
//
// 3) Install dependencies
//   go get -u github.com/stripe/stripe-go
//
// 4) Run the server on http://localhost:4242
//   go run server.go

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
	"github.com/stripe/stripe-go/v72/webhook"
	"io/ioutil"
	"lib_stripe/env"
	"log"
	"net/http"
	"os"
)

var (
	stripeKey     string
	webhookPath   string
	webhookSecret string

	callbackListenAddr string
	callbackServiceURL string
)

func init() {
	env.LoadEnv() // load env from `.env` file
	stripeKey = os.Getenv("STRIPE_PRIVATE_KEY")
	webhookPath = os.Getenv("STRIPE_WEBHOOK_PATH")
	webhookSecret = os.Getenv("STRIPE_WEBHOOK_SECRET")
	callbackListenAddr = os.Getenv("CALLBACK_LISTEN_ADDR")
	callbackServiceURL = os.Getenv("CALLBACK_SERVICE_URL")
}

func main() {
	router := gin.Default()

	router.Any("ping", ping())
	router.Any(webhookPath, handleWebhook())
	router.GET("/order/success", orderSuccess())
	router.GET("/order/cancel", orderCancel())
	router.Any("/mission/createCheckoutSession", createCheckoutSession())

	log.Printf("http server listening on %s", callbackListenAddr)
	err := router.Run(callbackListenAddr)
	if err != nil {
		log.Fatalln("http server failed: ", err)
	}
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func createCheckoutSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sc := &client.API{}
		sc.Init(stripeKey, nil)

		params := &stripe.CheckoutSessionParams{
			SuccessURL: stripe.String(callbackServiceURL + webhookPath),
			CancelURL:  stripe.String(callbackServiceURL + webhookPath),
			//CancelURL:  stripe.String(callbackServiceURL + webhookPath + "?session_id={CHECKOUT_SESSION_ID}"),
			Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				&stripe.CheckoutSessionLineItemParams{
					PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
						Currency: stripe.String("usd"),
						ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
							Name: stripe.String("T-shirt"),
						},
						UnitAmount: stripe.Int64(2000),
					},
					Quantity: stripe.Int64(1),
				},
			},
		}
		checkoutSession, err := sc.CheckoutSessions.New(params)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("chechout mission failed: %v\n", err),
			})
			return
		}

		c.Redirect(http.StatusSeeOther, checkoutSession.URL)
	}
}

func orderSuccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "order success!",
		})
	}
}

func orderCancel() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "order cancel!",
		})
	}
}

func handleWebhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"error": fmt.Sprintf("Error reading request body: %v\n", err),
			})
			return
		}

		var ev StripeEvent
		if err = json.Unmarshal(body, &ev); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("parse stripe event failed: %v\n", err),
			})
			return
		}

		log.Printf("webhook called with event: %s, %s", ev.Id, ev.Type)

		endpointSecret := webhookSecret
		// Pass the request body and Stripe-Signature header to ConstructEvent, along
		// with the webhook signing key.
		event, err := webhook.ConstructEvent(body, c.GetHeader("Stripe-Signature"), endpointSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Error verifying webhook signature: %v\n", err),
			})
			return
		}

		// Unmarshal the event data into an appropriate struct depending on its Type
		switch event.Type {
		case "payment_intent.succeeded":
			// Then define and call a function to handle the event payment_intent.succeeded
		// ... handle other event types
		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok!",
		})
	}
}

type StripeEvent struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	APIVersion string `json:"api_version"`
}
