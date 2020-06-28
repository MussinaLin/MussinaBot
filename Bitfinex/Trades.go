package Bitfinex

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"log"
	"math"
	"time"
)

func GetFRR(priorSecs int, frrBias float64) float64{
	log.Println(fmt.Sprintf("[GetFRR] priorSecs:%d, frrBias:%f", priorSecs, frrBias))
	trades := getTradeHistory(priorSecs)
	size := len(trades.Snapshot)
	log.Println("trades.Snapshot size:",size)
	var FRR float64 = 0
	for _, trade := range trades.Snapshot {
		FRR += trade.Price
	}
	FRR = FRR / float64(size)
	FRR = FRR * 100  // conv to %
	FRR = FRR + FRR * frrBias * 0.01
	FRR = math.Ceil(FRR*1000000) / 1000000
	log.Println("FRR:",FRR)
	return FRR
}

//func GetRealTimeApy(priorSecs int) float64{
//	log.Println(fmt.Sprintf("[GetRealTimeApy] priorSecs:%d", priorSecs))
//	trades := getTradeHistory(priorSecs)
//	size := len(trades.Snapshot)
//	log.Println("trades.Snapshot size:",size)
//	var RR float64 = 0
//	for _, trade := range trades.Snapshot {
//		RR += trade.Price
//	}
//	FRR = FRR / float64(size)
//	FRR = FRR * 100  // conv to %
//	FRR = FRR + FRR * frrBias * 0.01
//	FRR = math.Ceil(FRR*1000000) / 1000000
//	log.Println("FRR:",FRR)
//}

func getTradeHistory(priorSecs int) *bitfinex.TradeSnapshot{
	now := time.Now()
	millis := now.UnixNano() / 1000000
	prior := now.Add(time.Duration(-priorSecs) * time.Second)
	millisStart := prior.UnixNano() / 1000000
	start := bitfinex.Mts(millisStart)
	end := bitfinex.Mts(millis)

	trades, err := bfRestClient.Trades.PublicHistoryWithQuery("fUSD", start, end, bitfinex.QueryLimitMax, bitfinex.NewestFirst)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return trades
}
