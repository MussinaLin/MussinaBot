package api

import (
	"MussinaBot/Bitfinex"
	"MussinaBot/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ProvideFunds struct{
	Id          int64 `json:"id"`
	Symbol      string `json:"symbol"`
	UpdatedTime string `json:"updated_time"`
	Amount		float64 `json:"amount"`
	Rate		float64 `json:"rate"`
	Period		int64 `json:"period"`
	PositionPair string `json:"position_pair"`
}
func GetProvideFunds(w http.ResponseWriter, req *http.Request){
	log.Println("[GetProvideFunds...]")
	creditsSnap := Bitfinex.GetActivePositions()
	credits := creditsSnap.Snapshot
	log.Println("Provided funds size:", len(credits))
	provideFunds := make([]ProvideFunds, 0, len(credits))
	for _, credit := range credits {
		funds := ProvideFunds{
			Id:          credit.ID,
			Symbol:      credit.Symbol,
			UpdatedTime: cnvTimestamp2String(credit.MTSOpened),
			Amount:      credit.Amount,
			Rate:        credit.Rate,
			Period:      credit.Period,
			PositionPair:credit.PositionPair,
		}
		provideFunds = append(provideFunds, funds)
	}
	fmt.Fprintf(w, utils.CnvStruct2Json(provideFunds))
}

func cnvTimestamp2String(timestamp int64) string{
	timestamp = timestamp / 1000
	t := time.Unix(timestamp, 0)
	strDate := t.Format(time.RFC3339)
	return strDate
}
