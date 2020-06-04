package Bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
)

func IsPlatformWorking() bool{
	c := rest.NewClient()
	ok, _ := c.Platform.Status()
	if ok != true{
		return false
	}else {
		return true
	}
}
