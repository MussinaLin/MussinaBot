package main

import (
	"MussinaBot/Bitfinex"
	"MussinaBot/utils"
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var oneTimeFlag bool = true
var ordersNoti *[]bitfinex.Notification = nil
var ordersProvidedCount map[int64]int32
//var activeOrderSize = 0
var fundingNotLendCount = 0

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
			if fundingNotLendCount == 0{
				marginFundingLoan(cfg)
			}else{
				checkOrderStatus(cfg.OrdersNotLendTh)
			}
		}else{
			log.Println("Bitfinex is down...")
		}
	}
}

func startBitfinexWS(cfg *utils.Config){
	if oneTimeFlag == true{
		// start ws
		Bitfinex.StartAvaliableBalanceWS(cfg.ApiKey, cfg.ApiSecret,cfg.WsURL)
		oneTimeFlag = false
	}
}

func marginFundingLoan(cfg *utils.Config){
	log.Println("[marginFundingLoan]...")
	availBalance := Bitfinex.GetAvailableBalance()
	log.Println("available balance:", availBalance)
	if availBalance < cfg.MinLoan{
		return
	}

	FRR := Bitfinex.GetFRR(cfg.FrrCalculatePriorSecs, cfg.FrrBias)
	orders := Bitfinex.GenOrders(availBalance, cfg.MaxSingleOrderAmount, cfg.MinLoan, cfg.BalanceLeft)
	orders = Bitfinex.AssignRate(FRR, cfg.FrrIncreaseRate, orders)
	//orders = Bitfinex.ModifyPeriod(orders, cfg.FrrLoanMonthRate)
	log.Println(orders)

	ordersNoti = Bitfinex.SubmitOrders(orders)
	submittedOrderCount := len(*ordersNoti)
	//time.Sleep(300 * time.Millisecond)
	if submittedOrderCount > 0{
		log.Println(fmt.Sprintf("submit order count:[%d]", submittedOrderCount))
		fundingNotLendCount++
	}else{
		log.Println(fmt.Sprintf("[ERROR] submit order fail...count:[%d]", submittedOrderCount))
	}
	//initOrdersProvidedCount(ordersNoti)
}

func initOrdersProvidedCount(ordersNoti *[]bitfinex.Notification){
	log.Println("[initOrdersProvidedCount]...")
	ordersProvidedCount = make(map[int64]int32)
	for _, order := range *ordersNoti{
		ordersProvidedCount[order.MessageID] = 0
	}
}

func checkOrderStatus(notLendTh int){
	if fundingNotLendCount >= notLendTh{
		// cancel all order
	}

	if !isAllFundProvided(){
		fundingNotLendCount++
	}else{
		fundingNotLendCount = 0
	}
}

func isAllFundProvided() bool{
	size := Bitfinex.GetActiveOrdersSize()
	if size > 0{
		return false
	}else{
		return true
	}
}
