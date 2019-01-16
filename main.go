package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omise/omise-go"
	"github.com/yutify/omise-api/handler"
	"net/http"
)

const (
	Prefix = "/v1"
	ChargesPath = "/charges"
)

type Client struct {
	*omise.Client
}

func main() {
	router := mux.NewRouter()
	v1 := router.PathPrefix(Prefix).Subrouter()

	//	// Charge rooting
	v1.Path(ChargesPath).HandlerFunc(handler.CreateCharge).Methods(http.MethodPost)
	s := &http.Server{
		Addr: ":8080",
		Handler: router,
	}
	fmt.Println("=== start server ===")
	fmt.Println(s.ListenAndServe())
}