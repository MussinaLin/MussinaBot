package main

import (
	"MussinaBot/utils"
	"log"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil{
		log.Fatalln(err.Error())
	}

	cfg=cfg
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
