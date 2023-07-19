package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

func main() {
	stripe.Key = "sk_test_51NLn6uAMAaJjSlCRSuYdPj052VLmxwrjFVhNN78Qr0VoodDpLhwG3GD7MZPqFZ5kstWpU5VUwEuJzzTp3COdgSdv00nEWuz3tB"

	http.HandleFunc("/create-payment-intent", createPaymentIntent)
	http.HandleFunc("/confirm-payment-intent", confirmPaymentIntent)
	http.HandleFunc("/capture-payment-intent", capturePaymentIntent)

	log.Fatal(http.ListenAndServe("localhost:4242", nil))
}

func createPaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var request struct {
		PaymentMethodID string `json:"payment_method_id"`
	}
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1000), // Amount in cents
		Currency: stripe.String("usd"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		PaymentMethod: stripe.String(request.PaymentMethodID),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Failed to create Payment Intent: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := struct {
		ClientSecret string `json:"client_secret"`
	}{
		ClientSecret: pi.ClientSecret,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func confirmPaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var request struct {
		PaymentIntentID string `json:"payment_intent_id"`
	}
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	pi, err := paymentintent.Confirm(request.PaymentIntentID, nil)
	if err != nil {
		log.Printf("Failed to confirm Payment Intent: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payment Intent confirmed successfully: %s", pi.ID)
}

func capturePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var request struct {
		PaymentIntentID string `json:"payment_intent_id"`
	}
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	pi, err := paymentintent.Capture(request.PaymentIntentID, nil)
	if err != nil {
		log.Printf("Failed to capture Payment Intent: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payment Intent captured successfully: %s", pi.ID)
}
