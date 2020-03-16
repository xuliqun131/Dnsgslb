package handlers

import (
	"net"
	"fmt"
	"time"
	"os"
	"syscall"
	"Dnsgslb/pkg/errors"

)

var status = false
var tcperr string


func TcpCheck(ip string, port string, timeout int) bool {
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
		// 错误处理
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
	return status
}
