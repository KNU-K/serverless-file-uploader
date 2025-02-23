// hello.go
package main

import (
	"log"
	"net/http"
	"serverless-file-uploader/internal/handler"
)

func main(){
	r:= handler.RegisterRoutes()
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}