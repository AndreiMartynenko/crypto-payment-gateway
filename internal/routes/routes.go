package routes

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AndreiMartynenko/crypto-payment-gateway/internal/services/payment"
	"github.com/AndreiMartynenko/crypto-payment-gateway/internal/services/verification"
)

func HandleInitiatePayment(w http.ResponseWriter, r *http.Request) {
	// Read the request body into a byte slice
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Convert the byte slice into a bytes.Reader
	bodyReader := bytes.NewReader(bodyBytes)

	// Call the payment service with the bytes.Reader
	resp, err := payment.InitiatePayment(bodyReader)
	if err != nil {
		http.Error(w, "Error communicating with payment service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body and write it back to the client
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintln(w, string(responseBody))
}

func HandleVerifyTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction_id")
	if transactionID == "" {
		http.Error(w, "Transaction ID is required", http.StatusBadRequest)
		return
	}
	resp, err := verification.VerifyTransaction(transactionID)
	if err != nil {
		http.Error(w, "Error communicating with verification service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintln(w, string(body))
}
