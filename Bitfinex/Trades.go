package Bitfinex

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"log"
	"math"
	"time"
)

func GetFRR(priorMinutes int32, frrBias float64) float64{
	log.Println(fmt.Sprintf("[GetFRR] priorMinutes:%d, frrBias:%f", priorMinutes, frrBias))
	trades := getTradeHistory(priorMinutes)
	size := len(trades.Snapshot)
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

func getTradeHistory(priorMinutes int32) *bitfinex.TradeSnapshot{
	now := time.Now()
	millis := now.UnixNano() / 1000000
	prior := now.Add(time.Duration(-priorMinutes) * time.Minute)
	millisStart := prior.UnixNano() / 1000000
	start := bitfinex.Mts(millisStart)
	end := bitfinex.Mts(millis)

	trades, err := bitfinexClient.Trades.PublicHistoryWithQuery("fUSD", start, end, bitfinex.QueryLimitMax, bitfinex.NewestFirst)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return trades
}
