package exchange

import (
	"context"
	"strconv"
	"time"

	"github.com/fat/common/constant"
	"github.com/fat/common/crypto"
	"github.com/fat/models"
	"github.com/fat/repository"
	"github.com/fat/usecase/webhook"
)

type Indodax interface {
	Order(ctx context.Context, arg repository.CreateOrderParams) (*models.ResponseOrderIndodax, error)
	OrderHistory(ctx context.Context, pair string) (*models.ResponseOrderHistoryIndodax, error)
	MarketHistory(ctx context.Context, payload models.MarketHistoryPayload) ([]models.MarketHistoryIndodax, error)
	MarketPendingtHistory(ctx context.Context, symbolPending string) (*models.ResponseMarketPendingIndodax, error)
	Info(ctx context.Context) (*models.ResponseInfoIndodax, error)
}

type IndodaxCtx struct {
	config  *models.Config
	webhook webhook.IndodaxWebhook
}

func NewIndodax(conf *models.Config) Indodax {
	return &IndodaxCtx{
		config:  conf,
		webhook: webhook.NewIndodaxWebhook(conf),
	}
}

func (i *IndodaxCtx) Order(ctx context.Context, arg repository.CreateOrderParams) (*models.ResponseOrderIndodax, error) {
	payload := &crypto.PayloadSHA512{
		Pair:       arg.OrderCrypto,
		Type:       arg.OrderType,
		Price:      arg.OrderPrice,
		IDR:        arg.Idr,
		BTC:        arg.Btc,
		Method:     constant.INODAX_TRADE_METHOD,
		Timestamp:  time.Now().UnixNano() / int64(time.Millisecond),
		PrivateKey: i.config.Env.IndodaxPrivateKey,
	}

	return i.webhook.Order(payload)
}

func (i *IndodaxCtx) Info(ctx context.Context) (*models.ResponseInfoIndodax, error) {
	payload := &crypto.PayloadSHA512{
		Method:     constant.INODAX_GET_INFO_METHOD,
		Timestamp:  time.Now().UnixNano() / int64(time.Millisecond),
		RecvWindow: strconv.Itoa(int(time.Now().UnixNano() / int64(time.Millisecond))),
		PrivateKey: i.config.Env.IndodaxPrivateKey,
	}

	return i.webhook.GetInfo(payload)
}

func (i *IndodaxCtx) OrderHistory(ctx context.Context, pair string) (*models.ResponseOrderHistoryIndodax, error) {
	payload := &crypto.PayloadSHA512{
		Pair:       pair,
		Method:     constant.INODAX_ORDER_HISTORY_METHOD,
		Timestamp:  time.Now().Unix(),
		PrivateKey: i.config.Env.IndodaxPrivateKey,
	}

	return i.webhook.GetOrderHistory(payload)
}

func (i *IndodaxCtx) MarketHistory(ctx context.Context, payload models.MarketHistoryPayload) ([]models.MarketHistoryIndodax, error) {
	marketHistory, err := i.webhook.MarketHistory(&payload)
	if err != nil {
		return nil, err
	}

	return marketHistory, nil
}

func (i *IndodaxCtx) MarketPendingtHistory(ctx context.Context, symbolPending string) (*models.ResponseMarketPendingIndodax, error) {
	marketPendingHistory, err := i.webhook.MarkePendingtHistory(symbolPending)
	if err != nil {
		return nil, err
	}

	return marketPendingHistory, nil
}
