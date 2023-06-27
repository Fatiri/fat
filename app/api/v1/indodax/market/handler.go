package indodaxmarketv1

import (
	"net/http"
	"time"

	"github.com/Fatiri/fat/common/wrapper"
	"github.com/Fatiri/fat/models"
	"github.com/Fatiri/fat/usecase/exchange"
	"github.com/gin-gonic/gin"
)

type MarketHandler struct {
	config  *models.Config
	indodax exchange.Indodax
}

func NewMarketHandler(config *models.Config) *MarketHandler {
	return &MarketHandler{
		config:  config,
		indodax: exchange.NewIndodax(config),
	}
}

// MarketHistory godoc
// @Summary      Show an market history data
// @Description  get market history
// @Accept       json
// @Produce      json
// @Param        Authorization  header   string  true  "Authentication header"
// @Param        message        body      models.MarketHistoryPayload  true  "market history payload"
// @Success      200            {object}  models.MarketHistory
// @Failure      400            {object}  wrapper.Response
// @Failure      404            {object}  wrapper.Response
// @Failure      401            {object}  wrapper.Response
// @Router       /indodax/market_history [post]
func (oh *MarketHandler) MarketHistory(ctx *gin.Context) {
	var req models.MarketHistoryPayload
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(loc)
	timeStampOne := timeNow.Unix()
	timeStampTwo := timeNow.Add(time.Minute * -10080).Unix()

	req.From = timeStampTwo
	req.To = timeStampOne
	mh, err := oh.indodax.MarketHistory(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	ctx.JSON(http.StatusOK, mh)
}
