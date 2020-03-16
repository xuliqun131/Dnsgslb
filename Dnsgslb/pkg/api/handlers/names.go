package handlers

import (
	"net/http"
	"Dnsgslb/pkg/api/types"
	"log"
	"Dnsgslb/pkg/models"
	"fmt"
)

func ListName(w http.ResponseWriter, r *http.Request) {

}

func AddName(w http.ResponseWriter, r *http.Request) {
	res := types.Name{}
	_, err := Unmarshalname(&res, r)
	if err != nil {
		log.Fatal("AddName err is: ", err)
	}
	err = models.InsertNames("xu.com", res.Name)
	if err != nil {
		log.Fatal("Insert name err is: ", err)
	}
	fmt.Println("Insert name success !!!")
	err = models.InsertNameType(res.Name, "xu.com","A", 500, 500)
	if err != nil {
		log.Fatal("Insert NameType err is: ", err)
	}
	fmt.Println("Insert Nametype success")
}

func DeleteName(w http.ResponseWriter, r *http.Request) {

}

func UpdateName(w http.ResponseWriter, r *http.Request) {

}