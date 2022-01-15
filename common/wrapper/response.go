package wrapper

import (
	"fmt"
	"runtime"
	"strings"
)

type Response struct {
	Status   bool   `json:"status"`
	Message  string `json:"message"`
	Location string `json:"location,omitempty"`
}

func Error(err error, env string) Response {
	pc, fn, line, _ := runtime.Caller(1)
	fnSplit := strings.Split(fn, "/")
	if env == "release" {
		return Response{
			Status:  false,
			Message: err.Error(),
		}
	}
	return Response{
		Status:   false,
		Message:  err.Error(),
		Location: fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fnSplit[len(fnSplit)-1], line),
	}
}

func RouteNotFound() Response {
	return Response{
		Status:  false,
		Message: "Page not found",
	}
}

func Success(message string) Response {
	return Response{
		Status:  true,
		Message: message,
	}
}
