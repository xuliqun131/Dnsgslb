package router

import (
	"net/http"
	"github.com/mux"
	"Dnsgslb/pkg/api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var Routers []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/records", handlers.AddRecords).Methods("POST")
	router.HandleFunc("/records", handlers.ListRecords).Methods("GET")
	router.HandleFunc("/records/{args}", handlers.DeleteRecords).Methods("DELETE")
	router.HandleFunc("/records", handlers.UpdateRecords).Methods("PUT")
	router.HandleFunc("/sendicmp", handlers.Sendicmp).Methods("POST")
	router.HandleFunc("/sendtcp", handlers.Sendtcp).Methods("POST")
	router.HandleFunc("/sendhttp", handlers.Sendhttp).Methods("POST")

	//
	//for _, route := range Routers {
	//	var handler http.Handler
	//	handler = route.HandlerFunc
	//	router.
	//		Methods(route.Method).
	//		Path(route.Pattern).
	//		Name(route.Name).
	//		Handler(handler)
	//}
	return router
}

