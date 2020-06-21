package api

import "MussinaBot/Bitfinex"

type GeneralResp struct{
	TotalBalance float64
	AvailableBalance float64
}

func GetGeneralResp() *GeneralResp{
	resp := GeneralResp{}
	resp.TotalBalance = Bitfinex.GetTotalBalance()
	resp.AvailableBalance = Bitfinex.GetAvailableBalance()
	return &resp
}
