package handlers

import (
	"github/social-network/config"
	"github/social-network/db"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	dbmgr db.Datastore
)

// Init will initialize the handlers
func Init() {
	// build db manager
	dbmgr = db.NewClient()
	// Initialize the router and handlers
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/v1/account", CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/v1/account/login", LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/v1/account", ListAccounts).Methods(http.MethodGet)
	r.HandleFunc("/v1/account/{id}", UpdateAccount).Methods(http.MethodPatch)
	r.HandleFunc("/v1/account/{id}", DeleteAccount).Methods(http.MethodDelete)

	http.ListenAndServe(":"+config.Port, r)
}
