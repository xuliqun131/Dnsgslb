package handlers

import (
	"net/http"
	"fmt"
	_"github.com/mux"
	_"net"
	"io/ioutil"
	"Dnsgslb/pkg/api/types"
	"encoding/json"
	_"github.com/pkg/errors"
	"time"
	"strings"
	"Dnsgslb/pkg/errors"
)

var httperr string

func Sendhttp(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//link := vars["url"]
	body, _ := ioutil.ReadAll(r.Body)
	str := []byte(body)
	c := types.HttpCheck{}
	err := json.Unmarshal(str, &c)
	if err != nil {
		fmt.Println("json err is: ", err)
	}

	for range time.Tick(time.Duration(c.Interval) * time.Second) {
		go httpcheck(c.Url, c.Timeout)
		if httperr != "" {
			fmt.Println("main err is:", httperr)
		}
		fmt.Println("status is : ", status)
	}
}

func httpcheck(url string, timeout int)  {
	fmt.Println("url is: ", url, "is connecting, please wait……")
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		if status != false {
			status = false
		}
		if strings.Contains(err.Error(), "Client.Timeout exceeded") {
			httperr = errors.ErrHttpTimeout.Error()
		}else if strings.Contains(err.Error(), "target machine actively refused it") {
			httperr = errors.ErrHttpRefused.Error()
		}
	}else {
		if status != true {
			status = true
		}
		fmt.Println("connected success HTTP !!! status code is :", resp.StatusCode)
	}
}



//fmt.Println("url is: ", url, "is connecting, please wait……")
//reps, err := http.Get(url)
//
//if err != nil {
//	return err
//}else {
//	fmt.Println("status: ", reps.StatusCode)
//}
//defer reps.Body.Close()
//return nil