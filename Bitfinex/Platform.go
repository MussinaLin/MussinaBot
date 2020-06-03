package Bitfinex

import (
	"MussinaBot/utils"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"

)

var cfg *utils.Config

func SetConfig(_cfg *utils.Config){
	cfg = _cfg
}

func IsPlatformWorking() bool{
	c := rest.NewClientWithURL("https://api-pub.bitfinex.com/v2/").
		Credentials(cfg.ApiKey, cfg.ApiSecret)
	ok, _ := c.Platform.Status()
	if ok != true{
		return false
	}else {
		return true
	}
}
