package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"

	"github.com/gorilla/mux"
)

type Mock struct {
	*httptest.Server
}

func NewMock() *Mock {
	r := mux.NewRouter()
	registerRoutes(r)

	s := httptest.NewServer(r)
	return &Mock{s}
}

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/identity/v2/oauth2/token", accessTokenHandler)

	r.HandleFunc("/za/pb/v1/accounts", accountListHandler)

	r.HandleFunc("/za/pb/v1/accounts/{accountId}/balance",
		accountBalanceHandler)

	r.HandleFunc("/za/pb/v1/accounts/{accountId}/transactions",
		accountTransactionsHandler)
}

func readResponseFile(path string) ([]byte, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", abs)
	return ioutil.ReadFile(path)
}

func accessTokenHandler(w http.ResponseWriter, r *http.Request) {
	res, err := readResponseFile("testdata/access_token.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func accountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	res, err := readResponseFile("testdata/account_balance.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func accountListHandler(w http.ResponseWriter, r *http.Request) {
	res, err := readResponseFile("testdata/account_list.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func accountTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	res, err := readResponseFile("testdata/account_transactions.json")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
