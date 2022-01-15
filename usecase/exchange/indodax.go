package exchange

import (
	"strconv"
	"time"

	"github.com/FAT/common/constant"
	"github.com/FAT/common/crypto"
	"github.com/FAT/models"
	"github.com/FAT/repository"
	"github.com/FAT/usecase/webhook"
	"github.com/gin-gonic/gin"
)

type Indodax interface {
	Order(ctx *gin.Context, arg repository.CreateOrderParams) (repository.Order, error)
}

type IndodaxCtx struct {
	config  *models.Config
	webhook webhook.IndodaxWebhook
}

func NewIndodaxOrder(conf *models.Config) Indodax {
	return &IndodaxCtx{
		config:  conf,
		webhook: webhook.NewIndodaxWebhook(conf),
	}
}

func (i *IndodaxCtx) Order(ctx *gin.Context, arg repository.CreateOrderParams) (repository.Order, error) {
	payload := &crypto.PayloadSHA512{
		Pair:      arg.OrderCrypto,
		Type:      arg.OrderType,
		Price:     strconv.Itoa(int(arg.OrderPrice)),
		Method:    constant.INODAX_TRADE_METHOD,
		Timestamp: arg.CreatedAt.UnixNano() / int64(time.Millisecond),
	}

	responseOrder, err := i.webhook.Order(payload)
	if err != nil {
		return repository.Order{}, err
	}

	arg.OrderID = responseOrder.Return.OrderID
	
	return i.config.Storage.CreateOrder(ctx, arg)
}
