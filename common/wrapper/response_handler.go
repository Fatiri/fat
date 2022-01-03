package wrapper

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(err error) gin.H {
	pc, fn, line, _ := runtime.Caller(1)
	return gin.H{
		"status":        false,
		"message": err.Error(),
		"location":  fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line),
	}
}

func RouteNotFound() gin.H {
	return gin.H{
		"status":        false,
		"message": "Page not found",
	}
}

func SuccessHandler() gin.H {
	return gin.H{
		"status":  true,
		"message": "success",
	}
}
