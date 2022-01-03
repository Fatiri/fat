package wrapper

import (
	"fmt"
	"os"
	"runtime"
)

type MultiLanguages struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

func (e *MultiLanguages) Error() string {
	if e.EN != "" {
		return e.EN
	} else if e.ID != "" {
		return e.ID
	}
	return "something went wrong"
}

func NewResponseMultiLang(languages MultiLanguages) *MultiLanguages {
	return &languages
}

type Error struct {
	Desc        MultiLanguages `json:"desc"`
	Code        int            `json:"code"`
	Err         error          `json:"-"`
	ErrLocation string         `json:"errorLocation,omitempty"`
}

func NewError(code TypeError, err error) *Error {
	pc, fn, line, _ := runtime.Caller(1)

	var respErr MultiLanguages
	if errValue, isMatch := err.(*MultiLanguages); isMatch {
		if errValue != nil {
			respErr = *errValue
		} else {
			respErr = MultiLanguages{
				ID: err.Error(),
				EN: err.Error(),
			}
		}
	} else {
		respErr = MultiLanguages{
			ID: err.Error(),
			EN: err.Error(),
		}
	}
	envApp := os.Getenv("ENV_APP")
	if envApp == "production" {
		newResponseLog(err).Show()
		return &Error{Desc: respErr, Code: int(code)}
	}

	return &Error{Desc: respErr, Code: int(code), ErrLocation: fmt.Sprintf("%s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line)}
}

func (e *Error) Error() string {
	return e.Err.Error()
}
