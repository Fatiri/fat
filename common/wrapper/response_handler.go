package wrapper

import (
	"fmt"
	"runtime"
	"strings"
)

type MessageError struct {
	Status   bool   `json:"status"`
	Message  string `json:"message"`
	Location string `json:"location"`
}

type MessageSuccess struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func ErrorHandler(err error) MessageError {
	pc, fn, line, _ := runtime.Caller(1)
	fnSplit := strings.Split(fn,"/")
	return MessageError{
		Status:   false,
		Message:  err.Error(),
		Location: fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fnSplit[len(fnSplit)-1], line),
	}
}

func RouteNotFound() MessageError {
	pc, fn, line, _ := runtime.Caller(1)
	fnSplit := strings.Split(fn,"/")
	return MessageError{
		Status:   false,
		Message:  "Page not found",
		Location: fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fnSplit[len(fnSplit)-1], line),
	}
}

func SuccessHandler(message string) MessageSuccess {
	return MessageSuccess{
		Status:  true,
		Message: message,
	}
}
