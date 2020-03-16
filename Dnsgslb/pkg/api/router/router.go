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

	router.HandleFunc("/domain", handlers.ListDomain).Methods("GET")
	router.HandleFunc("/domain/{domain}", handlers.AddDomain).Methods("POST")
	router.HandleFunc("/domain/{domain}", handlers.DeleteDomain).Methods("DELETE")
	router.HandleFunc("/domain/{domain}", handlers.UpdateDoamin).Methods("PUT")

	router.HandleFunc("/content", handlers.ListContent).Methods("GET")
	router.HandleFunc("/content/{content}", handlers.AddContent).Methods("POST")
	router.HandleFunc("/content/{content}", handlers.DeleteContent).Methods("DELETE")
	router.HandleFunc("/content/{content}", handlers.UpdateContent).Methods("PUT")

	router.HandleFunc("/monitors", handlers.SelectMonitors).Methods("GET")
	router.HandleFunc("/monitor", handlers.ListMonitors).Methods("GET")
	router.HandleFunc("/monitor", handlers.AddMonitors).Methods("POST")
	router.HandleFunc("/monitor/{monitor}", handlers.DeleteMonitors).Methods("DELETE")
	router.HandleFunc("/monitor/{monitor}", handlers.UpdateMonitors).Methods("PUT")

	router.HandleFunc("/type", handlers.ListType).Methods("GET")
	router.HandleFunc("/type", handlers.AddType).Methods("POST")
	router.HandleFunc("/type/{type}", handlers.DeleteType).Methods("DELETE")
	router.HandleFunc("/type/{type}", handlers.UpdateType).Methods("PUT")

	router.HandleFunc("/name", handlers.ListName).Methods("GET")
	router.HandleFunc("/name", handlers.AddName).Methods("POST")
	router.HandleFunc("/name/{name}", handlers.DeleteName).Methods("DELETE")
	router.HandleFunc("/name/{name}", handlers.UpdateName).Methods("PUT")

	router.HandleFunc("/view", handlers.ListView).Methods("GET")
	router.HandleFunc("/view", handlers.AddView).Methods("POST")
	router.HandleFunc("/view/{view}", handlers.DeleteView).Methods("DELETE")
	router.HandleFunc("/view/{view}", handlers.UpdateView).Methods("PUT")

	router.HandleFunc("/records", handlers.AddRecords).Methods("POST")
	router.HandleFunc("/records", handlers.ListRecords).Methods("GET")
	router.HandleFunc("/records/{recordsId}", handlers.SelectRecords).Methods("GET")
	router.HandleFunc("/records/{recordsId}", handlers.DeleteRecords).Methods("DELETE")
	router.HandleFunc("/records/{recordsId}", handlers.UpdateRecords).Methods("PUT")


	return router
}

