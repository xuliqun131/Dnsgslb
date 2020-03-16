package handlers

import (
	"net/http"
	"Dnsgslb/pkg/api/types"
	"log"
	"Dnsgslb/pkg/models"
	"github.com/mux"
	"fmt"
	"encoding/json"
	"strconv"
)


func ListType(w http.ResponseWriter, r *http.Request) {
	var b json.RawMessage
	res := types.DnsType{}
	listtype, err := models.SelectType(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("listerr is", err)
	}

	for _, v := range listtype {
		res = v
		b, _ = json.Marshal(res)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
	}
}

func AddType(w http.ResponseWriter, r *http.Request) {
	res := types.DnsType{}
	_, err := Unmarshaltype(&res, r)
	if err != nil {
		log.Fatal("AddType err is: ", err)
	}
	err = models.InsertType(res.Type, res.Description)
	if err != nil {
		log.Fatal("Insert type err is: ", err)
	}
	fmt.Println("Insert type success!!!")
}

func DeleteType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["type"]
	typeid, err := strconv.Atoi(args)
	err = models.DelType(typeid)
	if err != nil {
		log.Fatal("Delete type err is: ", err)
	}
	fmt.Println("Delete type success!!!")
}

func UpdateType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["type"]
	typeid, err := strconv.Atoi(args)
	res := types.DnsType{}
	_, err = Unmarshaltype(&res, r)
	if err != nil {
		log.Fatal("UpdateType err is: ", err)
	}
	err = models.ChangeType(res.Type, res.Description, typeid)
	if err != nil {
		log.Fatal("Change type err is: ", err)
	}
	fmt.Println("Update type success!!!")
}