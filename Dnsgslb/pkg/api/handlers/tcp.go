package handlers

import (
	"net/http"
	"net"
	"fmt"
	"time"
	"io/ioutil"
	"os"
	"syscall"
	"encoding/json"
	"Dnsgslb/pkg/api/types"
	"Dnsgslb/pkg/errors"

)

var status = false
var tcperr string

func Sendtcp(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	str := []byte(body)
	c := types.TcpCheck{}
	err := json.Unmarshal(str, &c)
	if err != nil {
		fmt.Println("tcp json err is : ", err)
	}

	for range time.Tick(time.Duration(c.Interval) * time.Second) {
		go tcpcheck(c.IP, c.Port, c.Timeout)
		if tcperr != "" {
			fmt.Println("main err is:", tcperr)
		}
		fmt.Println("status is : ", status)
	}
}

func tcpcheck(ip string, port string, timeout int) {
	fmt.Println("addr", ip+":"+port, "is connecting, please wait……")
	_, err := net.DialTimeout("tcp4", ip+":"+port, time.Duration(timeout) * time.Second)
	if err != nil {
		if status != false {
			status = false
		}

		netErr, _ := err.(net.Error)
		if netErr.Timeout() {
			tcperr = errors.ErrTimout.Error()
		}

		opErr, _ := netErr.(*net.OpError)
		switch t := opErr.Err.(type) {
		case *net.DNSError :
			tcperr = errors.ErrDnserr.Error()
		case *os.SyscallError :
			tcperr = errors.ErrSyscall.Error()
			if errno, ok := t.Err.(syscall.Errno); ok {
				switch errno {
				case syscall.ECONNREFUSED :
					tcperr = errors.ErrRefuseConn.Error()
				case syscall.ETIMEDOUT :
					tcperr = errors.ErrTimout.Error()
				case syscall.EHOSTUNREACH :
					tcperr = errors.ErrHostUnreach.Error()
				}
			}
		}
	}else{
		if status != true {
			status = true
		}
		fmt.Println("connected success client !!! ")
	}
}

//func NetError(err error) bool {
//	netErr, ok := err.(net.Error)
//	if !ok {
//		return false
//	}
//
//	if netErr.Timeout() {
//		errors.ErrTimout.Error()
//		return true
//	}
//
//	opErr, ok := netErr.(*net.OpError)
//	if !ok {
//		return false
//	}
//
//	switch t := opErr.Err.(type) {
//	case *net.DNSError :
//		errors.ErrDnserr.Error()
//		return true
//	case *os.SyscallError :
//		errors.ErrSyscall.Error()
//		return  true
//		if errno, ok := t.Err.(syscall.Errno); ok {
//			switch errno {
//			case syscall.ECONNREFUSED :
//				errors.ErrRefuseConn.Error()
//				return true
//			case syscall.ETIMEDOUT :
//				errors.ErrTimout.Error()
//				return true
//			case syscall.EHOSTUNREACH :
//				errors.ErrHostUnreach.Error()
//				return true
//			}
//		}
//	}
//	return false
//}