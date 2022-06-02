package exercises

import (
	"encoding/json"
	"net/http"
)

type Transaction struct {
	Id uint64
	Code string
	Currency string
	Value float64
	Issuer string
	Receiver string
	CreatedAt string
}

func GetAllHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	transactions := []Transaction {
		{
			Id: 1,
        	Code: "BRPAY",
        	Currency: "BRL",
        	Value: 42.00,
        	Issuer: "pica-pau",
        	Receiver: "pé-de-pano",
        	CreatedAt: "2022-06-02",
		},
		{
			Id: 2,
        	Code: "BRPAY",
        	Currency: "BRL",
        	Value: 66.00,
        	Issuer: "pé-de-pano",
        	Receiver: "pica-pau",
        	CreatedAt: "2022-06-02",
		},
		{
			Id: 3,
        	Code: "BRPAY",
        	Currency: "BRL",
        	Value: 120.00,
        	Issuer: "pica-pau",
        	Receiver: "pé-de-pano",
        	CreatedAt: "2022-06-02",
		},
	}

	json.NewEncoder(w).Encode(transactions)
}