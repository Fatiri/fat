package crypto

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
)

type PayloadSHA512 struct {
	Method     string  `json:"method"`
	Timestamp  int64   `json:"timestamp"`
	RecvWindow string  `json:"recwindow"`
	Type       string  `json:"type"`
	Price      string  `json:"price"`
	IDR        string `json:"idr"`
	BTC        string  `json:"btc"`
	Pair       string  `json:"pair"`
	OrderID    int64   `json:"order_id"`
	Data       string  `json:"data"`
	PrivateKey string  `json:"private_key"`
}

func (p *PayloadSHA512) GenerateHMACSHA512() string {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha512.New, []byte(p.PrivateKey))
	h.Write([]byte(p.Data))
	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

func (p *PayloadSHA512) GenerateGetInfoIndodax() (string, string) {
	data := url.Values{}
	data.Set("method", p.Method)
	data.Set("timestamp", strconv.Itoa(int(p.Timestamp)))
	data.Set("recWindow", p.RecvWindow)
	p.Data = data.Encode()
	sha := p.GenerateHMACSHA512()

	return sha, data.Encode()
}

func (p *PayloadSHA512) GenerateOrderIndodax() (string, string) {
	data := url.Values{}
	data.Set("method", p.Method)
	data.Set("timestamp", strconv.Itoa(int(p.Timestamp)))
	data.Set("pair", p.Pair)
	data.Set("type", p.Type)
	data.Set("price", p.Price)

	switch strings.ToLower(p.Type) {
	case "buy":
		data.Set("idr", p.IDR)
	case "sell":
		data.Set("btc", p.BTC)
	}

	p.Data = data.Encode()
	sha := p.GenerateHMACSHA512()

	return sha, data.Encode()
}

func (p *PayloadSHA512) GenerateGetOrderIndodax() (string, string) {
	data := url.Values{}
	data.Set("method", p.Method)
	data.Set("timestamp", strconv.Itoa(int(p.Timestamp)))
	data.Set("pair", p.Pair)
	data.Set("order_id", strconv.Itoa(int(p.OrderID)))

	p.Data = data.Encode()
	sha := p.GenerateHMACSHA512()

	return sha, data.Encode()
}

func (p *PayloadSHA512) GenerateGetOrderHistoryIndodax() (string, string) {
	data := url.Values{}
	data.Set("method", p.Method)
	data.Set("pair", p.Pair)
	data.Set("nonce", strconv.Itoa(int(p.Timestamp)))

	p.Data = data.Encode()
	sha := p.GenerateHMACSHA512()

	return sha, data.Encode()
}

func (p *PayloadSHA512) GenerateMarketDataTokoCrypto() url.Values {
	data := url.Values{}
	data.Set("method", p.Method)
	data.Set("timestamp", strconv.Itoa(int(p.Timestamp)))
	data.Set("pair", p.Pair)
	data.Set("order_id", strconv.Itoa(int(p.OrderID)))

	p.Data = data.Encode()
	sha := p.GenerateHMACSHA512()
	data.Set("siganture", sha)

	return data
}
