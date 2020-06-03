package main

import (
	"MussinaBot/utils"
	"MussinaBot/Bitfinex"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil{
		log.Fatalln(err.Error())
	}

	cfg=cfg
	// periodic job
	tick := time.NewTicker(time.Second * 5)
	go scheduler(tick)

	Bitfinex.SetConfig(cfg)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	tick.Stop()
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
func scheduler(tick *time.Ticker) {
	for range tick.C {
		log.Println(Bitfinex.IsPlatformWorking())
	}
}