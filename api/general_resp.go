package api

import (
	"MussinaBot/Bitfinex"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type GeneralResp struct{
	TotalBalance float64 `json:"total_balance"`
	AvailableBalance float64 `json:"available_balance"`
}

func GetGeneralResp(w http.ResponseWriter, req *http.Request){
	resp := GeneralResp{}
	resp.TotalBalance = Bitfinex.GetTotalBalance()
	resp.AvailableBalance = Bitfinex.GetAvailableBalance()
	result, err := json.Marshal(resp)
	if err != nil {
		log.Println("[ERROR] GetGeneralResp error...", err)
	}
	fmt.Fprintf(w, string(result))
}
