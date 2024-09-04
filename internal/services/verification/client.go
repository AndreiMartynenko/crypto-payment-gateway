package verification

import (
	"net/http"
)

func VerifyTransaction(transactionID string) (*http.Response, error) {
	resp, err := http.Get("http://localhost:8082/verify-transaction?transaction_id=" + transactionID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
