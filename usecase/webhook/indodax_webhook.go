package webhook

import (
	"encoding/json"
	"net/http"

	"github.com/FAT/common/client"
	"github.com/FAT/common/crypto"
	"github.com/FAT/models"
)

type IndodaxWebhook interface {
	Order(payload *crypto.PayloadSHA512) (*models.ResponseOrderIndodax, error)
}

type IndodaxWebhookCtx struct {
	config *models.Config
}

func NewIndodaxWebhook(conf *models.Config) IndodaxWebhook {
	return &IndodaxWebhookCtx{
		config: conf,
	}
}

func (iw *IndodaxWebhookCtx) Order(payload *crypto.PayloadSHA512) (*models.ResponseOrderIndodax, error) {
	param := new(client.ParamaterHttpClient)
	param.Method = http.MethodPost
	param.URL = iw.config.Env.IndodaxPrivateURL

	sign, urlEncode := payload.GenerateOrderIndodax()

	param.Headers = client.RequestClientHeaderIndodax(sign, urlEncode)

	response, errHttpClient := client.HttpClientV2(param, urlEncode)

	responseByte, errReadResponse := client.ReadHttpResponse(response)
	if errReadResponse != nil {
		return nil, errReadResponse
	}

	responseOrder := new(models.ResponseOrderIndodax)

	errUnmarshal := json.Unmarshal(responseByte, &responseOrder)
	if errUnmarshal != nil || responseOrder.Success != 1 {
		return nil, errUnmarshal
	}

	return responseOrder, errHttpClient
}
