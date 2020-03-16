package handlers

import (
	"net/http"
	"strconv"
	"time"
	"log"
	"fmt"
	"encoding/json"
	"github.com/mux"
	_"github.com/joeig/go-powerdns"
	"Dnsgslb/pkg/api/types"
	"Dnsgslb/pkg/models"

)

var task = map[int64]chan struct{}{}

func SelectRecords(w http.ResponseWriter, r *http.Request) {
	res := types.Records{}
	vars := mux.Vars(r)
	args := vars["recordsId"]
	recordId, err := strconv.Atoi(args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = models.ListRecords(&res, recordId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("records is: ", res)
	var b json.RawMessage
	b, _ = json.Marshal(res)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)

}

func ListRecords(w http.ResponseWriter, r *http.Request) {
	res := types.Records{}
	listrecords, err := models.SelectRecords(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var b json.RawMessage
	for _, v := range listrecords {

		res = v
		b, _ = json.Marshal(res)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept,Authorization")
		w.Write(b)
	}

}

func AddRecords(w http.ResponseWriter, r *http.Request) {
	res := types.Records{}
	_, err := Unmarshalreocrd(&res, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("reve err is:" ,err)
	}

	fmt.Println("上线")
	err = Online(res.Domain, res.Name, res.Type, res.TTL, res.Content)
	if err != nil {
		fmt.Println("online err is: ", err)
	}

	var contentresult bool
	var nameresult bool
	var cmresult bool
	cmresult, err = models.SelectContentMonitor(res.Content, res.Monitors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("select contentmonitors err is: ", err)
	}
	var ntresult bool
	ntresult, err = models.SelectNameType(res.Name, res.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("select nametype err is: ", err)
	}
	if cmresult == false && ntresult == true {

		contentresult, err = models.SelectContent(res.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Select content is err is: ", err)
		}

		if contentresult == false {
			err = models.InsertContent(res.Content)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert content err is: ", err)
			}
			fmt.Println("Insert content success!!!")
		}

		err = models.InsertContentMonitor(res.Content, res.Monitors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Insert content_monitor err is: ", err)
		}

		res.Id, err = models.InsertRecords(res.Name, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Insert Records err is: ", err)
		}
		fmt.Println("AddRecord success1 !!!")

	}else if cmresult == true && ntresult == false {

		nameresult, err = models.SelectNames(res.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Select names err is: ", err)
		}

		if nameresult == false {
			err = models.InsertNames(res.Domain, res.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert names err is: ", err)
			}
		}

		err = models.InsertNameType(res.Name, res.Domain, res.Type, res.TTL, res.Persistence)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Inset nametype err is : ", err)
		}

		res.Id, err = models.InsertRecords(res.Name, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Insert Records err is: ", err)
		}
		fmt.Println("AddRecord success2 !!!")

	}else if cmresult == false && ntresult == false {

		contentresult, err = models.SelectContent(res.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Select content is err is: ", err)
		}

		if contentresult == false {
			err = models.InsertContent(res.Content)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert content err is: ", err)
			}
			fmt.Println("Insert content success!!!")
		}

		err = models.InsertContentMonitor(res.Content, res.Monitors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Insert content_monitor err is: ", err)
		}

		nameresult, err = models.SelectNames(res.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Select names err is: ", err)
		}

		if nameresult == false {
			err = models.InsertNames(res.Domain, res.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert names err is: ", err)
			}
		}

		err = models.InsertNameType(res.Name, res.Domain, res.Type, res.TTL, res.Persistence)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Inset nametype err is : ", err)
		}

		res.Id, err = models.InsertRecords(res.Name, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Insert Records err is: ", err)
		}
		fmt.Println("AddRecord success3 !!!")

	}else if cmresult == true && ntresult == true {

		res.Id, err = models.InsertRecords(res.Name, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Insert Records err is: ", err)
		}
		fmt.Println("AddRecord success4 !!!")
	}

	resmonitor := types.Monitors{}
	_, err = models.SelcetMonitor(res.Monitors, &resmonitor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("ListMonitors err is: ", err)
	}
	fmt.Println("Monitors is :", res.Monitors)
	fmt.Println("Interval is :", resmonitor.Interval)
	ch := make(chan struct{})
	task[res.Id] = ch
	fmt.Println("res.id: ", res.Id)

	ticker := time.NewTicker(time.Duration(resmonitor.Interval) * time.Second)

	var laststatus = false
	if resmonitor.Type == "icmp" {
		argsmap:=map[string]interface{}{}
		p:=NewPingOption()
		go func(){
			for {
				select {
				case <-ch:
					fmt.Println("recv stop single")
					return
				case <-ticker.C:
					resstatus := p.IcmpCheck(res.Content, argsmap)
					if resstatus != laststatus {
						err := models.ChangeStatus(resstatus, res.Id)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
						fmt.Println("Change status success ! status is ", resstatus)
						laststatus = resstatus
					}
					if resstatus == false {
						fmt.Println("下线")
						err := Offline(res.Content)
						if err != nil {
							fmt.Println("offline err is: ", err)
						}
					}else if resstatus == true {
						fmt.Println("上线")
						err = Online(res.Domain, res.Name, res.Type, res.TTL, res.Content)
						if err != nil {
							fmt.Println("online err is: ", err)
						}
					}
				}
			}
		}()





	}else if resmonitor.Type == "tcp" {
		fmt.Println("不创建")
		go func(){
			for{
				select {
				case <- ch:
					fmt.Println("recv stop single")
					return
				case <- ticker.C:
					resstatus := TcpCheck(res.Content, resmonitor.Port, resmonitor.Timeout)
					if resstatus != laststatus {
						err := models.ChangeStatus(resstatus, res.Id)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
						fmt.Println("Change status success ! status is ", resstatus)
						laststatus = resstatus
						if resstatus == false {
							fmt.Println("下线")
							err := Offline(res.Content)
							if err != nil {
								fmt.Println("offline err is: ", err)
							}
						}else if resstatus == true {
							fmt.Println("上线")
							err = Online(res.Domain, res.Name, res.Type, res.TTL, res.Content)
							if err != nil {
								fmt.Println("online err is: ", err)
							}
						}
					}
				}
			}
		}()

	}else if resmonitor.Type == "http" {
		fmt.Println("不创建")
		go func() {
			for {
				select {
				case <- ch:
					fmt.Println("recv stop single")
					return
				case <- ticker.C:
					resstatus := HttpCheck(res.Content, resmonitor.Timeout)
					if resstatus != laststatus {
						err := models.ChangeStatus(resstatus, res.Id)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
						}
						fmt.Println("Change status success ! status is, id is: ", resstatus, res.Id)
						laststatus = resstatus
						if resstatus == false {
							fmt.Println("下线")
							err := Offline(res.Content)
							if err != nil {
								fmt.Println("offline err is: ", err)
							}
						}else if resstatus == true {
							fmt.Println("上线")
							err = Online(res.Domain, res.Name, res.Type, res.TTL, res.Content)
							if err != nil {
								fmt.Println("online err is: ", err)
							}
						}
					}
				}
			}
		}()
	}
}



func DeleteRecords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	args := vars["recordsId"]
	recordId, err := strconv.Atoi(args)
	recordsId := int64(recordId)
	if err != nil {
		log.Fatal("strconv err is: ", err)
	}

	oldcontent, err := models.Selectcontent(recordId)
	if err != nil {
		fmt.Println("select old content err is: ", err)
	}
	err = Offline(oldcontent)
	if err != nil {
		fmt.Println("offline err is: ", err)
	}
	fmt.Println("delete old content is", oldcontent)

	err = models.DelRecords(recordsId)
	if err != nil {
		log.Fatal("deleterecord err is: ", err)
	}
	for k, ch := range task {
		if recordsId == k {
			close(ch)
		}
	}
}


func UpdateRecords(w http.ResponseWriter, r *http.Request) {
	var resmonitor types.Monitors
	var res types.Records
	vars := mux.Vars(r)
	args := vars["recordsId"]
	recordId, err := strconv.Atoi(args)
	recordsId := int64(recordId)
	fmt.Println("recordId is :", recordsId)
	if err != nil {
		log.Fatal("strconv err is: ", err)
	}

	_, err = Unmarshalreocrd(&res, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("UpdateRecords err is: ", err)
	}
	fmt.Println("recv records is", res)
	var cmresult bool
	cmresult, err = models.SelectContentMonitor(res.Content, res.Monitors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Select content err is: ", err)
	}
	fmt.Println("cmresult is: ", cmresult)
	var ntresult bool
	ntresult, err = models.SelectNameType(res.Name, res.Type)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("select name err is: ", err)
	}
	fmt.Println("ntresult is: ", ntresult)

	var contentresult bool
	contentresult, err = models.SelectContent(res.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Select content err is:", err)
	}
	fmt.Println("contentresult is:", contentresult)

	var nameresult bool
	nameresult ,err = models.SelectNames(res.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Select name err is:", err)
	}

	if cmresult == false && ntresult == true {
		// 如果已经没有content，则插入新content和新content_mointors
		if contentresult == false {
			err := models.InsertContent(res.Content)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert content err is: ", err)
			}
			fmt.Println("Insert content success !!")
			err = models.InsertContentMonitor(res.Content, res.Monitors)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert contentmonitor err is: ", err)
			}
		}
		// 如果已经有content，则修改其对应的检测器
		err = models.ChangeContentMonitor(res.Content, res.Monitors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Change contentmonitor err is: ", err)
		}
		fmt.Println("Change contentmonitor success !!")

		err = models.ChangeRecords(res.Name, res.Domain, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View, recordsId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("ChangeRecords err is: ", err)
		}
		fmt.Println("changeRcords success1 !!")

	}else if cmresult == true && ntresult == false {

		if nameresult == false{
			err := models.InsertNames(res.Domain, res.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Inset name err is: ", err)
			}
			fmt.Println("Insert name success !!")
			// 如果没有则添加
			err = models.InsertNameType(res.Name, res.Domain, res.Type, res.TTL, res.Persistence)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert nametype err is: ", err)
			}
		}
		// 有则修改
		err = models.ChangeNameType(res.Type, res.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("change nametype err is: ", err)
		}
		fmt.Println("change nametype success !!")

		err = models.ChangeRecords(res.Name, res.Domain, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View, recordsId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("ChangeRecords err is: ", err)
		}
		fmt.Println("changeRcords2 success !!")

	}else if cmresult == false && ntresult == false {

		if contentresult == false {
			err := models.InsertContent(res.Content)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert content err is: ", err)
			}
			fmt.Println("Insert content success !!")
			err = models.InsertContentMonitor(res.Content, res.Monitors)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert contentmonitor err is: ", err)
			}
			fmt.Println("Insert contentmonitor success !!")
		}
		err = models.ChangeContentMonitor(res.Content, res.Monitors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Change contentmonitor err is: ", err)
		}
		fmt.Println("Change contentmonitor success !!")


		if nameresult == false {
			err = models.InsertNames(res.Domain, res.Name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Inset name err is: ", err)
			}
			fmt.Println("Insert name success !!")
			err = models.InsertNameType(res.Name, res.Domain, res.Type, res.TTL, res.Persistence)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Insert nametype err is: ", err)
			}
			fmt.Println("Insert nametype success !!")
		}

		err = models.ChangeNameType(res.Type, res.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("change nametype err is: ", err)
		}
		fmt.Println("change nametype success !!")

		err = models.ChangeRecords(res.Name, res.Domain, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View, recordsId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("ChangeRecords err is: ", err)
		}
		fmt.Println("changeRcords3 success !!")

	}else if cmresult == true && ntresult == true {
		err = models.ChangeRecords(res.Name, res.Domain, res.Type, res.Content, res.TTL, res.Disable, res.Weight, res.Monitors, res.View, recordsId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("ChangeRecords err is: ", err)
		}
		fmt.Println("changeRcords4 success !!")
	}
	fmt.Println("UpdateRecords success!!!")


	_, err = models.SelcetMonitor(res.Monitors, &resmonitor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("ListMonitors err is: ", err)
	}
	fmt.Println("recordsId is ", recordsId)


	var laststatus = false
	ticker := time.NewTicker(time.Duration(resmonitor.Interval) * time.Second)
	if resmonitor.Type == "icmp" {
		//关闭旧通道，旧goroutine，开新goroutine
		for k, lastch := range task {
			if k == recordsId {
				close(lastch)
			}
		}
		ch := make(chan struct{})
		task[recordsId] = ch
		//ping 检测函数
		argsmap:=map[string]interface{}{}
		p:=NewPingOption()
		go func(){
			for {
				select {
				case <-ch:
					fmt.Println("recv stop single")
					return
				case <-ticker.C:
					resstatus := p.IcmpCheck(res.Content, argsmap)
					if resstatus != laststatus {
						err := models.ChangeStatus(resstatus, recordsId)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							fmt.Println("change status err is: ", err)
						}
						fmt.Println("Change status success ! status is, id is: ", resstatus, recordsId)
						if resstatus == true {
							fmt.Println("上线,content", res.Content)

							err = Online(res.Domain, res.Name, res.Type, res.TTL, res.Content)
							if err != nil {
								fmt.Println("online err is: ", err)
							}
						}

						laststatus = resstatus
					}
					if resstatus == false {
						fmt.Println("下线,content", res.Content)

						err = Offline(res.Content)
						if err != nil {
							fmt.Println("offline err is: ", err)
						}
					}
					err := models.ChangeStatus(resstatus, recordsId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						fmt.Println("change status err is: ", err)
					}
				}
			}
		}()
		fmt.Println("Last status is ", laststatus)
	}else if resmonitor.Type == "tcp" {
		for k, lastch := range task {
			if k == recordsId {
				close(lastch)
			}
		}
		ch := make(chan struct{})
		task[recordsId] = ch
		go func(){
			for{
				select {
				case <- ch:
					fmt.Println("recv stop single")
					return
				case <- ticker.C:
					resstatus := TcpCheck(res.Content, resmonitor.Port, resmonitor.Timeout)
					if resstatus != laststatus {
						err := models.ChangeStatus(resstatus, recordsId)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							fmt.Println("change status err is: ", err)
						}
						fmt.Println("Change status success ! status is, id is: ", resstatus, recordsId)
						if resstatus == true {
							fmt.Println("上线,content", res.Content)

							err = Online(res.Domain, res.Name, res.Type, res.TTL, res.Content)
							if err != nil {
								fmt.Println("online err is: ", err)
							}
						}
						laststatus = resstatus
					}
					if resstatus == false {
						fmt.Println("下线,content", res.Content)

						err = Offline(res.Content)
						if err != nil {
							fmt.Println("offline err is: ", err)
						}
					}
					err := models.ChangeStatus(resstatus, recordsId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						fmt.Println("change status err is: ", err)
					}
				}
			}
		}()

	}else if resmonitor.Type == "http" {
		fmt.Println("上线")
		for k, lastch := range task {
			if k == recordsId {
				close(lastch)
			}
		}
		ch := make(chan struct{})
		task[recordsId] = ch
		go func() {
			for {
				select {
				case <- ch:
					fmt.Println("recv stop single")
					return
				case <- ticker.C:
					resstatus := HttpCheck(res.Content, resmonitor.Timeout)
					if resstatus != laststatus {
						err := models.ChangeStatus(resstatus, recordsId)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							fmt.Println("change status err is: ", err)
						}
						fmt.Println("Change status success ! status is, id is: ", resstatus, recordsId)
						if resstatus == true {
							fmt.Println("上线,content", res.Content)

							err = Online(res.Domain, res.Name, res.Type, res.TTL, res.Content)
							if err != nil {
								fmt.Println("online err is: ", err)
							}
						}
						laststatus = resstatus
					}
					if resstatus == false {
						fmt.Println("下线,content", res.Content)

						err = Offline(res.Content)
						if err != nil {
							fmt.Println("offline err is: ", err)
						}
					}
					err := models.ChangeStatus(resstatus, recordsId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						fmt.Println("change status err is: ", err)
					}
				}
			}
		}()
	}
}

