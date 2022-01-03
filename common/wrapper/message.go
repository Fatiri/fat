package wrapper

import (
	"encoding/json"
	"net/http"
)

type MultilangMessage struct {
	In string `json:"english"`
	Id string `json:"Bahasa"`
}

// MultilangErr struct
type Multilang struct {
	Stats bool             `json:"status"`
	Msg   MultilangMessage `json:"message"`
}

// MultiStringError use to response error with multi languange
type MultiStringError Multilang

// NewMultiStringError construct use to return the error
func NewMultiStringError(sts bool, msg MultilangMessage) *MultiStringError {
	return &MultiStringError{
		Stats: sts,
		Msg:   msg,
	}
}

func (c *MultiStringError) Status() bool {
	return c.Stats
}

func (c *MultiStringError) Message() MultilangMessage {
	return c.Msg
}

func (c *MultiStringError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// Message function response
func ResponseHeader(stats bool, message MultilangMessage) Multilang {
	return Multilang{
		Stats: stats,
		Msg:   message,
	}
}

// WriteToHeader function response to json
func WriteToHeader(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)
}
