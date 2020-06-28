package api

import (
	"MussinaBot/Bitfinex"
	"MussinaBot/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type GeneralResp struct{
	TotalBalance float64 `json:"total_balance"`
	AvailableBalance float64 `json:"available_balance"`
	Rate float64 `json:"rate"`
}

type EarnedInterest struct {

}

func GetGeneralResp(w http.ResponseWriter, req *http.Request){
	log.Println("[GetGeneralResp...]")
	resp := GeneralResp{}
	resp.TotalBalance = Bitfinex.GetTotalBalance()
	resp.AvailableBalance = Bitfinex.GetAvailableBalance()
	resp.Rate = Bitfinex.GetFRR(60, 0)
	fmt.Fprintf(w, utils.CnvStruct2Json(resp))
}

func GetEarnedInterest(w http.ResponseWriter, req *http.Request){
	log.Println("[GetEarnedInterest...]")
	now := time.Now().Unix()
	s := int64(86400 * 2450) // the day before 2450 days. 2450 is max_limit of ledgers count.
	start := now - s
	now = now * 1000
	start = start * 1000
	log.Println(fmt.Sprintf("start:%d end:%d", start, now))
	ledgers := Bitfinex.GetLedgers("USD", start, now)
	var c int32 = 0
	for i, ledger := range *ledgers{
		if strings.Contains(ledger.Description, "Margin Funding Payment"){
			log.Println(fmt.Sprintf("[%d]  Date:%s, Amount:%f, Balance:%f, Currency:%s",
				i, utils.CnvTimestamp2String(ledger.MTS), ledger.Amount, ledger.Balance, ledger.Currency))
			c++
		}
	}
	log.Println("c:",c)
}
