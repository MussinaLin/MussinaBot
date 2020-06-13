package Bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
)
var apiKey string
var apiSecret string
var url string

var bfRestClient *rest.Client

type WalletType int

const (
	Exchange WalletType = iota
	Funding
)

func (w WalletType) String() string {
	return []string{"exchange", "funding"}[w]
}

func SetConfig(_apiKey string, _apiSecret string, _url string){
	apiKey = _apiKey
	apiSecret = _apiSecret
	url = _url
	bfRestClient = rest.NewClientWithURL(url).Credentials(apiKey, apiSecret)
}
