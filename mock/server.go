package mock

import (
	"net/http"
	"net/http/httptest"
)

type Server struct {
	*httptest.Server
}

func NewServer(h http.Handler) *Server {
	return &Server{httptest.NewServer(h)}
}

var transactionsResponse = []byte(`
{
    "data": {
        "transactions": [
            {
                "accountId": "172878438321553632224",
                "type": "DEBIT",
                "status": "POSTED",
                "description": "MONTHLY SERVICE CHARGE",
                "cardNumber": "",
                "postingDate": "2020-06-11",
                "valueDate": "2020-06-10",
                "actionDate": "2020-06-18",
                "transactionDate": "2020-06-10",
                "amount": 535
            },
            {
                "accountId": "172878438321553632224",
                "type": "CREDIT",
                "status": "POSTED",
                "description": "CREDIT INTEREST",
                "cardNumber": "",
                "postingDate": "2020-06-11",
                "valueDate": "2020-06-10",
                "actionDate": "2020-06-18",
                "transactionDate": "2020-06-10",
                "amount": 31.09
            }
        ]
    },
    "links": {
        "self": "https://openapi.investec.com/za/pb/v1/accounts/{accountId}/transactions"
    },
    "meta": {
        "totalPages": 1
    }
}
`)

func (srv *Server) GetAccountTransactions(w http.ResponseWriter,
	r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(transactionsResponse)
}
