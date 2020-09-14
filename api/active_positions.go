package api

import (
	"MussinaBot/Bitfinex"
	"MussinaBot/httputil"
	"MussinaBot/utils"
	"fmt"
	"log"
	"math"
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
	ExpireAt 	float64 `json:"expire_at"`
	PositionPair string `json:"position_pair"`
}
func GetProvideFunds(w http.ResponseWriter, req *http.Request){
	log.Println("[GetProvideFunds...]")
	httputil.EnableCors(&w)
	creditsSnap := Bitfinex.GetActivePositions()
	credits := creditsSnap.Snapshot
	log.Println("Provided funds size:", len(credits))
	provideFunds := make([]ProvideFunds, 0, len(credits))
	for _, credit := range credits {
		funds := ProvideFunds{
			Id:          credit.ID,
			Symbol:      credit.Symbol,
			UpdatedTime: utils.CnvTimestamp2String(credit.MTSOpened),
			Amount:      credit.Amount,
			Rate:        utils.CnvDailyRate2AnnualRate(credit.Rate),
			Period:      credit.Period,
			ExpireAt:    getHowFarBeforeExpireTime(credit.MTSOpened, credit.Period),
			PositionPair:credit.PositionPair,
		}
		provideFunds = append(provideFunds, funds)
	}
	fmt.Fprintf(w, utils.CnvStruct2Json(provideFunds))
}

func getHowFarBeforeExpireTime(createMTS int64, period int64) float64{
	createMTS = createMTS / 1000
	periodMTS := 86400 * period
	expireTime := periodMTS + createMTS
	now := time.Now().Unix()
	result := expireTime - now
	resultF := float64(result) / float64(3600)
	resultF = math.Round(resultF*10) / 10
	return resultF
}
