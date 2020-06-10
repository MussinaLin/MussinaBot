package main

import (
	"MussinaBot/Bitfinex"
	"MussinaBot/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var oneTimeFlag bool = true

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil{
		log.Fatalln(err.Error())
	}

	// periodic job
	tick := time.NewTicker(time.Second * 6)
	go scheduler(tick, cfg)

	Bitfinex.SetConfig(cfg.ApiKey, cfg.ApiSecret, cfg.PubEndpoint)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	tick.Stop()
	log.Println("Stop timer...")
	Bitfinex.CloseWS()
	//c := rest.NewClientWithURL("https://api.bitfinex.com/v2/").Credentials(key, secret)
	//
	//snapHist, err := c.Funding.OfferHistory("fUSD")
	//if err != nil {
	//	panic(err)
	//}
	//for _, item := range snapHist.Snapshot {
	//	fmt.Println(item)
	//}
}
func scheduler(tick *time.Ticker, cfg *utils.Config) {
	for range tick.C {
		if Bitfinex.IsPlatformWorking(){
			startBitfinexWS(cfg)
			// loan algorithm
			log.Println("Bitfinex is up...")
			marginFundingLoan(cfg)
		}else{
			log.Println("Bitfinex is down...")
		}
	}
}

func startBitfinexWS(cfg *utils.Config){
	if oneTimeFlag == true{
		// start ws
		Bitfinex.StartAvaliableBalanceWS(cfg.ApiKey, cfg.ApiSecret,"")
		oneTimeFlag = false
	}
}

func marginFundingLoan(cfg *utils.Config){
	availBalance := Bitfinex.GetAvailableBalance()
	log.Println("available balance:", availBalance)
	if availBalance < cfg.MinLoan{
		//return
	}
	FRR := Bitfinex.GetFRR(10, cfg.FrrBias)
	orders := Bitfinex.GenOrders(availBalance, cfg.MaxSingleOrderAmount, cfg.MinLoan, cfg.BalanceLeft)
	orders = Bitfinex.AssignRate(FRR, cfg.FrrIncreaseRate, orders)
	fmt.Println(orders)
}
