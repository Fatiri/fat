package webhook

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fat/common/client"
	"github.com/fat/common/crypto"
	"github.com/fat/models"
)

type IndodaxWebhook interface {
	Order(payload *crypto.PayloadSHA512) (*models.ResponseOrderIndodax, error)
	GetOrder(payload *crypto.PayloadSHA512) (*models.ResponseGetOrderIndodax, error)
	GetOrderHistory(payload *crypto.PayloadSHA512) (*models.ResponseOrderHistoryIndodax, error)
	GetInfo(payload *crypto.PayloadSHA512) (*models.ResponseInfoIndodax, error)
	MarketHistory(payload *models.MarketHistoryPayload) ([]models.MarketHistoryIndodax, error)
	MarkePendingtHistory(symbol string) (*models.ResponseMarketPendingIndodax, error)
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

	param.Headers = client.RequestClientHeaderIndodax(sign, urlEncode, iw.config.Env.IndodaxPublicKey)

	response, errHttpClient := client.HttpClientV2(param, urlEncode)
	if errHttpClient != nil {
		return nil, errHttpClient
	}

	responseByte, errReadResponse := client.ReadHttpResponse(response)
	if errReadResponse != nil {
		return nil, errReadResponse
	}

	responseOrder := new(models.ResponseOrderIndodax)
	responseError := new(models.ResponseErrorIndodax)

	errUnmarshal := json.Unmarshal(responseByte, &responseOrder)
	if errUnmarshal != nil || responseOrder.Success != 1 {
		json.Unmarshal(responseByte, &responseError)
		return nil, errors.New(responseError.Error)
	}

	return responseOrder, nil
}

func (iw *IndodaxWebhookCtx) GetOrder(payload *crypto.PayloadSHA512) (*models.ResponseGetOrderIndodax, error) {
	param := new(client.ParamaterHttpClient)
	param.Method = http.MethodPost
	param.URL = iw.config.Env.IndodaxPrivateURL

	sign, urlEncode := payload.GenerateOrderIndodax()

	param.Headers = client.RequestClientHeaderIndodax(sign, urlEncode, iw.config.Env.IndodaxPublicKey)

	response, errHttpClient := client.HttpClientV2(param, urlEncode)
	if errHttpClient != nil {
		return nil, errHttpClient
	}

	responseByte, errReadResponse := client.ReadHttpResponse(response)
	if errReadResponse != nil {
		return nil, errReadResponse
	}

	responseGetOrder := new(models.ResponseGetOrderIndodax)
	responseError := new(models.ResponseErrorIndodax)

	errUnmarshal := json.Unmarshal(responseByte, &responseGetOrder)
	if errUnmarshal != nil || responseGetOrder.Success != 1 {
		json.Unmarshal(responseByte, &responseError)
		return nil, errors.New(responseError.Error)
	}

	return responseGetOrder, nil
}

func (iw *IndodaxWebhookCtx) GetOrderHistory(payload *crypto.PayloadSHA512) (*models.ResponseOrderHistoryIndodax, error) {
	param := new(client.ParamaterHttpClient)
	param.Method = http.MethodPost
	param.URL = iw.config.Env.IndodaxPrivateURL

	sign, urlEncode := payload.GenerateGetOrderHistoryIndodax()

	param.Headers = client.RequestClientHeaderIndodax(sign, urlEncode, iw.config.Env.IndodaxPublicKey)

	response, errHttpClient := client.HttpClientV2(param, urlEncode)
	if errHttpClient != nil {
		return nil, errHttpClient
	}

	responseByte, errReadResponse := client.ReadHttpResponse(response)
	if errReadResponse != nil {
		return nil, errReadResponse
	}

	responseOrderHistory := new(models.ResponseOrderHistoryIndodax)
	responseError := new(models.ResponseErrorIndodax)

	errUnmarshal := json.Unmarshal(responseByte, &responseOrderHistory)
	if errUnmarshal != nil || responseOrderHistory.Success != 1 {
		json.Unmarshal(responseByte, &responseError)
		return nil, errors.New(responseError.Error)
	}

	return responseOrderHistory, nil
}

func (iw *IndodaxWebhookCtx) GetInfo(payload *crypto.PayloadSHA512) (*models.ResponseInfoIndodax, error) {
	param := new(client.ParamaterHttpClient)
	param.Method = http.MethodPost
	param.URL = iw.config.Env.IndodaxPrivateURL

	sign, urlEncode := payload.GenerateOrderIndodax()

	param.Headers = client.RequestClientHeaderIndodax(sign, urlEncode, iw.config.Env.IndodaxPublicKey)

	response, errHttpClient := client.HttpClientV2(param, urlEncode)
	if errHttpClient != nil {
		return nil, errHttpClient
	}

	responseByte, errReadResponse := client.ReadHttpResponse(response)
	if errReadResponse != nil {
		return nil, errReadResponse
	}

	responseInfo := new(models.ResponseInfoIndodax)
	responseError := new(models.ResponseErrorIndodax)

	errUnmarshal := json.Unmarshal(responseByte, &responseInfo)
	if errUnmarshal != nil || responseInfo.Success != 1 {
		json.Unmarshal(responseByte, &responseError)
		return nil, errors.New(responseError.Error)
	}

	return responseInfo, nil
}

func (iw *IndodaxWebhookCtx) MarketHistory(payload *models.MarketHistoryPayload) ([]models.MarketHistoryIndodax, error) {
	param := new(client.ParamaterHttpClient)
	param.Method = http.MethodGet
	param.URL = iw.config.Env.IndodaxMarketHistoryURL
	param.Query = client.RequestClientHeaderIndodaxV2(payload)
	response, errHttpClient := client.HttpClient(param)
	if errHttpClient != nil {
		return nil, errHttpClient
	}

	responseByte, errReadResponse := client.ReadHttpResponse(response)
	if errReadResponse != nil {
		return nil, errReadResponse
	}

	var responseInfo []models.MarketHistoryIndodax

	errUnmarshal := json.Unmarshal(responseByte, &responseInfo)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return responseInfo, nil
}

func (iw *IndodaxWebhookCtx) MarkePendingtHistory(symbol string) (*models.ResponseMarketPendingIndodax, error) {
	param := new(client.ParamaterHttpClient)
	param.Method = http.MethodGet
	param.URL = iw.config.Env.IndodaxMarketPendingURL + symbol

	response, errHttpClient := client.HttpClient(param)
	if errHttpClient != nil {
		return nil, errHttpClient
	}

	responseByte, errReadResponse := client.ReadHttpResponse(response)
	if errReadResponse != nil {
		return nil, errReadResponse
	}

	marketPending := new(models.ResponseMarketPendingIndodax)

	errUnmarshal := json.Unmarshal(responseByte, &marketPending)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return marketPending, nil
}
