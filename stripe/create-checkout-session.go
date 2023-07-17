package main

import (
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func main() {
	//stripe.Key = "sk_test_51NLn6uAMAaJjSlCRqP73JO9EnAyybtGXZUUP3g1h61F3sxKai6GAJdu0NxDeHofudqYFFW6DmiHXX7BMoBXq5cIU00bEQrKs2r"

	http.HandleFunc("/create-checkout-session", handleCreateCheckoutSession)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	// Create a new Checkout Session
	/*params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Product Name"),
					},
					UnitAmount: stripe.Int64(1000), // amount in cents
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)), */

		stripe.Key = "sk_test_51NLn6uAMAaJjSlCRqP73JO9EnAyybtGXZUUP3g1h61F3sxKai6GAJdu0NxDeHofudqYFFW6DmiHXX7BMoBXq5cIU00bEQrKs2r"

		params := &stripe.CheckoutSessionParams{
  			LineItems: []*stripe.CheckoutSessionLineItemParams{
    		{
      			Price: stripe.String("price_H5ggYwtDq4fbrJ"),
      			Quantity: stripe.Int64(2),
    		},
			},
  			Mode: stripe.String("payment"),
  			SuccessURL: stripe.String("https://google.com"),
		}
		s, _ := session.New(params)
		// SuccessURL: stripe.String("https://google.com"),
		// CancelURL:  stripe.String("https://droidsans.com"),

	}

	/*session, err := session.New(params)
	if err != nil {
		log.Printf("Error creating checkout session: %s", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}*/

	// Return the Checkout Session ID in the response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"sessionId": "` + session.ID + `"}`))
}
