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
	OrderPrice  string `json:"order_price"`
	OrderType   string `json:"order_type"`
	OrderCrypto string `json:"order_crypto"`
}

func (o OrderPayload) Validate() error {
	var err error
	listCryptoType := []CryptoType{BTC}
	for _, Ctype := range listCryptoType {
		if o.OrderCrypto != string(Ctype) {
			err = fmt.Errorf("order_crypto must be one in list : %v", listCryptoType)
		}
	}

	if o.OrderPrice == "" {
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

type ResponseErrorIndodax struct {
	Success   int64  `json:"success"`
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
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

type ResponseGetOrderIndodax struct {
	Success int64                         `json:"success"`
	Return  ResponseRetrunGetOrderIndodax `json:"return"`
}

type ResponseRetrunGetOrderIndodax struct {
	Order ResponseReturnGetOrderIndodax `json:"order"`
}

type ResponseReturnGetOrderIndodax struct {
	Price      string `json:"price"`
	Type       string `json:"type"`
	OrderRP    string `json:"order_rp"`
	RemainRP   string `json:"remain_rp"`
	SubmitTime string `json:"submit_time"`
	FinishTime string `json:"finish_time"`
	OrderID    string `json:"order_id"`
	Status     string `json:"status"`
}

type ResponseOrderHistoryIndodax struct {
	Success int64                             `json:"success"`
	Return  ResponseRetrunOrderHistoryIndodax `json:"return"`
}

type ResponseRetrunOrderHistoryIndodax struct {
	Orders []ResponseReturnOrderHistoryIndodax `json:"orders"`
}

type ResponseReturnOrderHistoryIndodax struct {
	OrderID    string `json:"order_id"`
	Type       string `json:"type"`
	Price      string `json:"price"`
	SubmitTime string `json:"submit_time"`
	FinishTime string `json:"finish_time"`
	Status     string `json:"status"`
	OrderIdr   string `json:"order_idr"`
	RemainIdr  string `json:"remain_idr"`
	OrderBtc   string `json:"order_btc"`
	RemainBtc  string `json:"remain_btc"`
}

type ResponseInfoIndodax struct {
	Success int64                   `json:"success"`
	Return  ResponseInfoIndodaxData `json:"return"`
}

type ResponseInfoIndodaxData struct {
	ServerTime         int64      `json:"server_time"`
	UserID             string     `json:"user_id"`
	Name               string     `json:"name"`
	Email              string     `json:"email"`
	ProfilePicture     string     `json:"profile_picture"`
	VerificationStatus string     `json:"verification_status"`
	GauthEnable        bool       `json:"gauth_enable"`
	Balance            InfoDetail `json:"balance"`
	BalanceHold        InfoDetail `json:"balance_hold"`
	Address            InfoDetail `json:"address"`
}

type ResponseMarketPendingIndodax struct {
	BuyOrders  []MarketPendingBuyIndodax  `json:"buy_orders"`
	SellOrders []MarketPendingSellIndodax `json:"sell_orders"`
}

type MarketPendingBuyIndodax struct {
	Price string `json:"price"`
	SumRP string `json:"sum_rp"`
}

type MarketPendingSellIndodax struct {
	Price  string `json:"price"`
	SumBTC string `json:"sum_btc"`
	SumETH string `json:"sum_eth"`
	SumALT string `json:"sum_alt"`
}

type InfoDetail struct {
	IDR   float64 `json:"idr"`
	BTC   string  `json:"btc"`
	ALT   string  `json:"alt"`
	ABYSS string  `json:"abyss"`
}

type MarketHistoryPayload struct {
	Symbol    string
	TimeFrame string
	From      int64
	To        int64
}

type MarketHistoryIndodax struct {
	Time   int64  `json:"time"`
	Open   int64  `json:"open"`
	High   int64  `json:"high"`
	Low    int64  `json:"low"`
	Close  int64  `json:"close"`
	Volume string `json:"volume"`
}

type MarketHistory struct {
	Open  []float64 `json:"open"`
	High  []float64 `json:"high"`
	Low   []float64 `json:"low"`
	Close []float64 `json:"close"`
}
