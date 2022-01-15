package models

import (
	"fmt"
)

type CryptoType string
type OrderType string

const (
	BTC  CryptoType = "btc_idr"
	Sell OrderType  = "sell"
	Buy  OrderType  = "buy"
)

type OrderPayload struct {
	OrderPrice  float64 `json:"order_price"`
	OrderType   string  `json:"order_type"`
	OrderCrypto string  `json:"order_crypto"`
}

func (o OrderPayload) Validate() error {
	var err error
	listCryptoType := []CryptoType{BTC}
	for _, Ctype := range listCryptoType {
		if o.OrderCrypto != string(Ctype) {
			err = fmt.Errorf("order_crypto must be one in list : %v", listCryptoType)
		}
	}

	var infFloat float64
	if o.OrderPrice == 0 && infFloat == float64(int(o.OrderPrice)) {
		err = fmt.Errorf("order_price can't be 0 & must be numeric")
	}

	listOrderType := []OrderType{Buy, Sell}
	for _, Otype := range listOrderType {
		if o.OrderType != string(Otype) {
			err = fmt.Errorf("order_type must be one in list : %v", listOrderType)
		}
	}
	
	return err
}

type ResponseOrderIndodax struct {
	Success int64                      `json:"success"`
	Return  ResponseReturnOrderIndodax `json:"return"`
}

type ResponseReturnOrderIndodax struct {
	ReceiveBtc string `json:"receive_btc"`
	SpendRP    int64  `json:"spend_rp"`
	Fee        int64  `json:"fee"`
	RemainRP   int64  `json:"remain_rp"`
	OrderID    int64  `json:"order_id"`
}
