package client

import "net/url"

func QueryParamTokoCryptoKlines() url.Values {
	urlValue := url.Values{}
	urlValue.Set("symbol", "ZILBIDR")
	urlValue.Set("interval", "1h")
	urlValue.Set("limit", "5")

	return urlValue
}
