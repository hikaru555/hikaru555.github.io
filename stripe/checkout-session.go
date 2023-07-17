package main

import (
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func main() {
	// This is your test secret API key.
	stripe.Key = "sk_test_51NLn6uAMAaJjSlCRqP73JO9EnAyybtGXZUUP3g1h61F3sxKai6GAJdu0NxDeHofudqYFFW6DmiHXX7BMoBXq5cIU00bEQrKs2r"

	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/stripe/checkout-session.go", createCheckoutSession)
	addr := "hikaru555.github.io/stripe"
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func createCheckoutSession(w http.ResponseWriter, r *http.Request) {
	domain := "https://hikaru555.github.io/stripe"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("{{PRICE_ID}}"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
