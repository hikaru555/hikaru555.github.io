package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

// This example sets up an endpoint using the Echo framework.
// Watch this video to get started: https://youtu.be/ePmEVBu8w6Y.

func main() {
	stripe.Key = "sk_test_51NLn6uAMAaJjSlCRqP73JO9EnAyybtGXZUUP3g1h61F3sxKai6GAJdu0NxDeHofudqYFFW6DmiHXX7BMoBXq5cIU00bEQrKs2r"

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/create-checkout-session", createCheckoutSession)

	e.Logger.Fatal(e.Start("hikaru555.github.io/stripe"))
}

func createCheckoutSession(c echo.Context) (err error) {
	params := &stripe.CheckoutSessionParams{
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
		SuccessURL: stripe.String("http://google.com"),
		CancelURL:  stripe.String("http://droidsans.com"),
	}

	s, _ := session.New(params)

	if err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, s.URL)
}
