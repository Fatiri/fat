package constant

// URL path
var (
	INDODAX_RoutePathTicker          = "/api/ticker/"
	INDODAX_RoutePathMarketHistory   = "/tradingview/history_v2"
	INDODAX_RoutePathPrivateAPI      = "/tapi"
	INDODAX_RoutePathGetOrderPending = "/api/webdata/BTCIDR"
)

// Paramaeter URL
var (
	INDODAX_SymbolParam    = "symbol"
	INDODAX_FromParam      = "from"
	INDODAX_ToParam        = "to"
	INDODAX_TimeFrameParam = "tf"
)

// Env key name
var (
	INDODAX_KeyEnvMainDomain = "INDODAX_MAIN_DOMAIN"
	INDODAX_KeyCryptoTrading = "INDODAX_CRYPTO_CURRENCY"
	INDODAX_KeyCryptoType    = "INDODAX_CRYPTO_TYPE"
	INDODAX_KeyEnvPeriodTime = "INDODAX_PERIODE_TIME"
	INDODAX_KeyEnvTimeFrame  = "INDODAX_TIME_FRAME"
	INDODAX_KeyEnvValueBuy   = "INDODAX_VALUE_BUY"
	INDODAX_KeyEnvValueSell  = "INDODAX_VALUE_SELL"
	INDODAX_KeyEnvPair       = "INDODAX_PAIR"
)

// indodax variable
var (
	INDODAX_BTCIDR          = "btcidr"
	INDODAX_BTCIDRSnakeCase = "btc_idr"
	Indodax                 = "INDODAX"
	INDODAX_MethodeGetInfo  = "getInfo"
	INDODAX_MethodeGetOrder = "getOrder"
	INDODAX_MethodeTrade    = "trade"
)

// public
var (
	TimeZone     = "Asia/Jakarta"
	Buy          = "buy"
	Sell         = "sell"
	Open         = "open"
	Hold         = "hold"
	IDRLowerCase = "idr"
	Filled       = "filled"
	Cancelled    = "cancelled"
	Low          = "low"
	High         = "high"
)
