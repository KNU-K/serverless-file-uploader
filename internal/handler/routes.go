package handler

import (
	"serverless-file-uploader/internal/handler/ping"
	"serverless-file-uploader/internal/handler/upload"

	"github.com/gorilla/mux"
)
func RegisterRoutes() * mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/ping", ping.Handler)
	api.HandleFunc("/upload", upload.Handler)
	return r
}