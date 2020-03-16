package handlers

import (
	"net/http"
	"github.com/mux"
	"Dnsgslb/pkg/models"
	"log"
	"Dnsgslb/pkg/api/types"
	"Dnsgslb/pkg/errors"
	"fmt"
	"encoding/json"
	"strconv"
)


func SelectMonitors(w http.ResponseWriter, r *http.Request) {
	var b json.RawMessage
	res := types.Monitors{}
	listmonitor, err := models.ListallMonitor(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	}
	for _, v := range listmonitor {
		res = v
		b, _ := json.Marshal(res)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
	}
	fmt.Println("monitoralljson is", string(b))
}

func ListMonitors(w http.ResponseWriter, r *http.Request) {
	var b json.RawMessage
	res := types.Monitors{}
	listmonitor, err := models.ListMonitor(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, v := range listmonitor {
		res = v
		b, _ := json.Marshal(res)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(b)
	}
	fmt.Println("monitorjson is", string(b))
}

func AddMonitors(w http.ResponseWriter, r *http.Request) {
	res := types.Monitors{}
	_, err := Unmarshalmonitor(&res, r)
	if err != nil {
		log.Fatal("AddMonitors err is: ", err)
	}
	err = models.InsertMonitor(res.Monitor, res.Type, res.Content, res.Port, res.Interval, res.Timeout)
	if err != nil {
		log.Fatal("Insert Monitor err is:", err)
	}
	fmt.Println("Insert monitor success!!!")

}

func DeleteMonitors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["monitor"]
	monitorid, err := strconv.Atoi(args)
	err = models.DelMonitor(monitorid)
	if err != nil {
		log.Fatal("DeleteMonitor err is: ", err)
	}
	fmt.Println("Deletemonitor success!!!")
}

func UpdateMonitors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["monitor"]
	monitorid, err := strconv.Atoi(args)
	res := types.Monitors{}
	_, err = Unmarshalmonitor(&res, r)
	if err != nil {
		errors.ErrJsonUnmarshal.Error()
	}

	err = models.ChangeMonitor(res.Monitor, res.Type, res.Content, res.Port, res.Interval, res.Timeout, monitorid)
	if err != nil {
		log.Fatal("Update monitor err is: ", err)
	}
	fmt.Println("Updatemonitor success!!!")
}