package handlers

import (
	"net/http"
	"github.com/mux"
	"Dnsgslb/pkg/models"
	"log"
	"Dnsgslb/pkg/api/types"
	"Dnsgslb/pkg/errors"
	"fmt"
	"github.com/joeig/go-powerdns"
	"encoding/json"
	"strconv"
)

func ConnPowerdns() *powerdns.Client {
	pdns := powerdns.NewClient("http://localhost:8081", "localhost", map[string]string{"X-API-Key": "apipw"}, nil)
	return pdns
}

func ListDomain(w http.ResponseWriter, r *http.Request) {

	var b json.RawMessage
	res := types.Domains{}
	listdomain, err := models.SelectDomain(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, v := range listdomain {
		res = v
		b, _ = json.Marshal(res)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
	}

}


func AddDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["domain"]
	err := models.InsertDomain(args)
	if err != nil {
		fmt.Println("AddDomain err is: ", err)
	}
	pdns := ConnPowerdns()
	_, err = pdns.Zones.AddNative(args, false, "", false, "", "", true, []string{"localhost."})
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Println("Insert domain success!!!")
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["domain"]
	domainid, err := strconv.Atoi(args)
	domainname, err := models.Selectdomainname(domainid)
	err = models.DelDomain(domainid)
	if err != nil {
		fmt.Println("DeleteDomain err is: ", err)
	}

	if err != nil {
		fmt.Println("Select domainname err is:", err)
	}
	pdns := ConnPowerdns()
	err = pdns.Zones.Delete(domainname)
	if err != nil {
		fmt.Println("Delete zone err is:", err)
	}
	fmt.Println("DeleteDomain success!!!")
}

func UpdateDoamin(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	args := vars["domain"]
	domainid, err := strconv.Atoi(args)
	res := types.Domains{}
	_, err = Unmarshaldomain(&res, r)
	if err != nil {
		errors.ErrJsonUnmarshal.Error()
	}
	// 删除zone
	domainname, err := models.Selectdomainname(domainid)
	if err != nil {
		fmt.Println("Select domainname err is:", err)
	}
	pdns := ConnPowerdns()
	err = pdns.Zones.Delete(domainname)
	if err != nil {
		fmt.Println("Delete zone err is:", err)
	}

	_, err = pdns.Zones.AddNative(res.Domain, false, "", false, "", "", true, []string{"localhost."})
	if err != nil {
		fmt.Println("Change zone err is:", err)
	}

	err = models.ChangeDoamin(res.Domain, domainid)
	if err != nil {
		log.Fatal("UpdateDomain err is: ", err)
	}
	fmt.Println("update domain success!!!")
}