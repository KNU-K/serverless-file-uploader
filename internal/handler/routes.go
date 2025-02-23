package handler

import (
	"github.com/gorilla/mux"
)
func RegisterRoutes() * mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/ping", handlePing)
	return r
}