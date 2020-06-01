package main


import (
	"fmt"
	//"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/rest"
	//"os"
)

func main() {
	//key := os.Getenv("BFX_KEY")
	//secret := os.Getenv("BFX_SECRET")
	key := ""
	secret := ""
	c := rest.NewClientWithURL("https://api.bitfinex.com/v2/").Credentials(key, secret)

	snapHist, err := c.Funding.OfferHistory("fUSD")
	if err != nil {
		panic(err)
	}
	for _, item := range snapHist.Snapshot {
		fmt.Println(item)
	}
}
