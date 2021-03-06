package main

import (
	"MussinaBot/Bitfinex"
	"MussinaBot/httputil"
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
	Bitfinex.SetConfig(cfg.ApiKey, cfg.ApiSecret, cfg.PubEndpoint)

	// process first
	processBitfinexLoan(cfg)

	// periodic job
	tick := time.NewTicker(time.Second * 60)
	go scheduler(tick, cfg)

	// start httputil server
	httputil.StartHttpServer()

	waitAndClose(tick)
}
func scheduler(tick *time.Ticker, cfg *utils.Config) {
	for range tick.C {
		processBitfinexLoan(cfg)
	}
}

func processBitfinexLoan(cfg *utils.Config){
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

func startBitfinexWS(cfg *utils.Config){
	if oneTimeFlag == true{
		// start ws
		Bitfinex.StartAvaliableBalanceWS(cfg.ApiKey, cfg.ApiSecret,cfg.WsURL)
		oneTimeFlag = false
		time.Sleep(5 * time.Second) // sleep to wait ws update balance
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
	orders = Bitfinex.ModifyPeriod(orders, cfg.FrrLoanMonthRate)
	//log.Println(orders)

	ordersNoti = Bitfinex.SubmitOrders(orders)
	submittedOrderCount := len(*ordersNoti)
	if submittedOrderCount > 0{
		log.Println(fmt.Sprintf("submit order count:[%d]", submittedOrderCount))
		fundingNotLendCount++
	}else{
		log.Println(fmt.Sprintf("[ERROR] submit order fail...count:[%d]", submittedOrderCount))
	}
}


func checkOrderStatus(notLendTh int){
	log.Println(fmt.Sprintf("[checkOrderStatus] fundingNotLendCount:[%d]",fundingNotLendCount))
	// get all orders
	orders := Bitfinex.GetActiveOrders()

	if fundingNotLendCount >= notLendTh{
		log.Println("cancel all orders...")
		fundingNotLendCount = 0

		// cancel all order
		if orders != nil{
			Bitfinex.CancelAllOrders(orders)
		}
		return
	}

	if !isAllFundProvided(orders){
		log.Println("Funds are not all provided...")
		fundingNotLendCount++
	}else{
		log.Println("===== Funds are all provided =====")
		fundingNotLendCount = 0
	}
}

func isAllFundProvided(orders *[]*bitfinex.Offer) bool{
	if orders != nil{
		size := len(*orders)
		if size > 0{
			return false
		}else{
			return true
		}
	}else{
		return true
	}
}

func waitAndClose(tick *time.Ticker){
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	tick.Stop()
	log.Println("Stop timer...")
	Bitfinex.CloseWS()
}

