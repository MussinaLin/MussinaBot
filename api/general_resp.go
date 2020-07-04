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
	RateDaily float64 `json:"rate_daily"`
}

type InterestSummary struct {
	TotalEarned float64 `json:"total_earned"`
	Apy float64 `json:"apy"`
	DailyInterests []DailyInterest `json:"daily_interests"`
}

type DailyInterest struct {
	Date string `json:"date"`
	Amount float64 `json:"amount"`
	Rate float64 `json:"rate"`
}

func GetGeneralResp(w http.ResponseWriter, req *http.Request){
	log.Println("[GetGeneralResp...]")
	resp := GeneralResp{}
	resp.TotalBalance = Bitfinex.GetTotalBalance()
	resp.AvailableBalance = Bitfinex.GetAvailableBalance()
	resp.RateDaily = Bitfinex.GetFRR(300, 0)
	resp.Rate = utils.CnvDailyRate2AnnualRate(resp.RateDaily)
	fmt.Fprintf(w, utils.CnvStruct2Json(resp))
}

func GetEarnedInterest(w http.ResponseWriter, req *http.Request){
	log.Println("[GetEarnedInterest...]")
	now := time.Now().Unix()
	s := int64(86400 * 500) // the day before 500 days. 500 is max_limit of ledgers count.
	start := now - s
	now = now * 1000
	start = start * 1000
	ledgers := Bitfinex.GetLedgers("USD", start, now)
	
	interests := make([]DailyInterest, 0, 50)
	var totalEarned float64 = 0.0
	var days int64 = 0
	for _, ledger := range *ledgers{
		if strings.Contains(ledger.Description, "Margin Funding Payment"){
			dailyEarn := DailyInterest{
				Date:   utils.CnvTimestamp2String(ledger.MTS),
				Amount: ledger.Amount,
				Rate:   utils.GetApyFromDailyInterest(Bitfinex.GetTotalBalance(), ledger.Amount),
			}
			interests = append(interests, dailyEarn)
			totalEarned += ledger.Amount
			days++
		}
	}
	resp := InterestSummary{
		TotalEarned:    utils.RoundFloat(totalEarned),
		Apy:            getApyFromTotalEarned(totalEarned, Bitfinex.GetTotalBalance(), days),
		DailyInterests: interests,
	}
	fmt.Fprintf(w, utils.CnvStruct2Json(resp))
}

func getApyFromTotalEarned(totalEarned float64, balance float64, days int64) float64{
	if balance == 0{
		log.Println("[ERROR] balance is 0")
		balance = 6738.56
	}
	return utils.RoundFloat( (totalEarned / balance) * 365 / float64(days))
}
