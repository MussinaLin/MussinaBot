package api

import (
	"MussinaBot/Bitfinex"
	"MussinaBot/utils"
	"fmt"
	"log"
	"net/http"
)

type GeneralResp struct{
	TotalBalance float64 `json:"total_balance"`
	AvailableBalance float64 `json:"available_balance"`
}

func GetGeneralResp(w http.ResponseWriter, req *http.Request){
	log.Println("[GetGeneralResp...]")
	resp := GeneralResp{}
	resp.TotalBalance = Bitfinex.GetTotalBalance()
	resp.AvailableBalance = Bitfinex.GetAvailableBalance()
	fmt.Fprintf(w, utils.CnvStruct2Json(resp))
}
