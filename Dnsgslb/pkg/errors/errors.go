package errors

import (
	"fmt"


)

//func StandardError

type StandardError struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}


var (
	ErrRefuseConn 		=	StandardError{1001, "Connected Refuse"}
	ErrTimout 			=	StandardError{1002, "ClientTimeout"}
	ErrUnknowHost		=	StandardError{1003, "UnknowHost"}
	ErrBadRequest		= 	StandardError{1004, "BadRequest"}
	ErrUnknowError		= 	StandardError{1005, "UnknowError"}
	ErrDnserr			= 	StandardError{1006, "DnsError"}
	ErrSyscall			=	StandardError{1007, "SyscallError"}
	ErrHostUnreach		=	StandardError{1008, "HostUnreach"}
	ErrHttpTimeout		= 	StandardError{1009, "HttpConnectedTimeout"}
	ErrHttpRefused		=	StandardError{1010, "Target host refused it"}
	ErrJsonUnmarshal	=	StandardError{1011, "JsonUnmarshal fail"}
)


func (err StandardError) Error() string{
	return fmt.Sprintf("errorCode is %d, errorMsg is %s", err.ErrorCode, err.ErrorMsg)
}

