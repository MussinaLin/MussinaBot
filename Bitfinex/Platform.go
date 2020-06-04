package Bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
)

var apiKey string
var apiSecret string
var url string
func SetConfig(_apiKey string, _apiSecret string, _url string){
	apiKey = _apiKey
	apiSecret = _apiSecret
	url = _url
}

func IsPlatformWorking() bool{
	c := rest.NewClientWithURL(url).Credentials(apiKey, apiSecret)
	ok, _ := c.Platform.Status()
	if ok != true{
		return false
	}else {
		return true
	}
}
