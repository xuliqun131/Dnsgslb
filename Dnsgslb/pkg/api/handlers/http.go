package handlers

import (
	"net/http"
	"fmt"
	_"github.com/mux"
	_"net"
	_"github.com/pkg/errors"
	"time"
	"strings"
	"Dnsgslb/pkg/errors"
)

var httperr string


func HttpCheck(url string, timeout int) bool {
	fmt.Println("url is: ", url, "is connecting, please wait……")
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		if status != false {
			status = false
		}
		// 错误处理
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
	return status
}


