package handlers

import (
	"net/http"
	"os/exec"
	"github.com/mux"
	"Dnsgslb/pkg/models"
	"github.com/caicloud/nirvana/log"
	"fmt"
	"Dnsgslb/pkg/api/types"
	"Dnsgslb/pkg/errors"

)

func ListContent(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("/bin/bash", "-c", "pdnsutil add-record xu.com liqun A 500 "+"192.0.2.102")
	if err := cmd.Start(); err != nil {
		fmt.Printf("Erro:The command is err", err)
	}
}

func AddContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["content"]
	err := models.InsertContent(args)
	if err != nil {
		log.Fatal("AddContent err is: ", err)
	}
	err = models.InsertContentMonitor(args, "icmp")
	if err != nil {
		log.Fatal("insert contentmonitor err is :", err)
	}
	fmt.Println("Insert Content sucCcess!!!")
}

func DeleteContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["content"]
	err := models.DelContent(args)
	if err != nil {
		log.Fatal("DeleteContent err is: ", err)
	}
	fmt.Println("Delete Content success!!!")
}

func UpdateContent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//args为原来的值，res.content为新值
	args := vars["content"]
	res := types.ContentType{}
	_, err := Unmarshalcontent(&res, r)
	if err != nil {
		errors.ErrJsonUnmarshal.Error()
	}
	fmt.Println("args is : ", args)
	fmt.Println("content is: ", res.Content)
	err = models.ChangeContent(res.Content, args)
	if err != nil {
		log.Fatal("UpdateContent err is: ", err)
	}
	fmt.Println("update content success!!!")
}
