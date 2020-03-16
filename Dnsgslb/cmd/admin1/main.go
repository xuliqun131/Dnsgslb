package main

import (
	"net/http"
	"Dnsgslb/pkg/api/router"
	"github.com/gorilla/handlers"
)

func main() {
	router := router.NewRouter()
	//log.Fatal(http.ListenAndServe(":8080", router))
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router))
}


