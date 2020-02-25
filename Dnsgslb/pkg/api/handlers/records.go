package handlers

import (
	"net/http"
	"Dnsgslb/pkg/models"
	"fmt"
	"io/ioutil"
	"Dnsgslb/pkg/api/types"
	"encoding/json"
)


func AddRecords(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	str := []byte(body)
	res := types.Records{}
	err := json.Unmarshal(str, &res)
	if err != nil {
		fmt.Println("Addrecords json ummarshal err is :", err)
	}
	if res.Monitors == "icmp" {

	}
	err = models.InsertRecords(res.Domain, res.Name, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View)
	if err != nil {
		fmt.Println("Add records err is: ", err)
	}
}

func ListRecords(w http.ResponseWriter, r *http.Request) {

}

func DeleteRecords(w http.ResponseWriter, r *http.Request) {

}

func UpdateRecords(w http.ResponseWriter, r *http.Request) {

}
