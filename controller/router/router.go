package router

import (
	"github.com/gorilla/mux"
	"github.com/yutify/gengo-api/controller/handler"
	"net/http"
)

const (
	Prefix      = "/v1"
	ChargesPath = "/charges"
	UsersPath   = "/users"
)

func Route(h handler.ApiHandler) http.Handler {
	router := mux.NewRouter()
	v1 := router.PathPrefix(Prefix).Subrouter()
	// Charge rooting
	v1.Path(ChargesPath).HandlerFunc(h.CreateCharge).Methods(http.MethodPost)
	// User rooting
	v1.Path(UsersPath).HandlerFunc(h.CreateUser).Methods(http.MethodPost)
	return router
}
