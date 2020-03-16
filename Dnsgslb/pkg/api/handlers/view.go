package handlers

import (
	"net/http"
	"Dnsgslb/pkg/api/types"
	"log"
	"github.com/mux"
	"Dnsgslb/pkg/models"
	"fmt"
	"encoding/json"
)

func ListView(w http.ResponseWriter, r *http.Request) {
	var b json.RawMessage
	res := types.Views{}
	listview, err := models.ListView(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, v := range listview {
		res = v
		b, _ := json.Marshal(res)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
	}

}

func AddView(w http.ResponseWriter, r *http.Request) {
	res := types.Views{}
	_, err := Unmarshalview(&res, r)
	if err != nil {
		log.Fatal("AddView err is: ", err)
	}
	err = models.InsertView(res.View, res.Rule)
	if err != nil {
		log.Fatal("Insert view err is: ", err)
	}
	fmt.Println("Insert view success !!!")
}

func DeleteView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["view"]
	err := models.DelView(args)
	if err != nil {
		log.Fatal("Delete view err is: ", err)
	}
	fmt.Println("Delete view success !!!")
}

func UpdateView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["view"]
	res := types.Views{}
	_, err := Unmarshalview(&res, r)
	if err != nil {
		log.Fatal("Update view err is: ", err)
	}
	err = models.ChangeView(res.View, res.Rule, args)
	if err != nil {
		log.Fatal("Update view err is: ", err)
	}
	fmt.Println("Update view success !!!")
}