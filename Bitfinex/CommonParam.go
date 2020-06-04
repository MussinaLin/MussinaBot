package Bitfinex

var apiKey string
var apiSecret string
var url string

func SetConfig(_apiKey string, _apiSecret string, _url string){
	apiKey = _apiKey
	apiSecret = _apiSecret
	url = _url
}
