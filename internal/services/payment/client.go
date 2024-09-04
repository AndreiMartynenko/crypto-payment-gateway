package payment

import (
	"bytes"
	"net/http"
)

func InitiatePayment(body *bytes.Reader) (*http.Response, error) {
	resp, err := http.Post("http://localhost:8081/initiate-payment", "application/json", body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
