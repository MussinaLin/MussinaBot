package Bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"log"
	"time"
)

func GetAvaliableBalanceWS(apiKey string, apiSecret string, uri string) float64{
	p := websocket.NewDefaultParameters()

	//p.URL = uri
	c := websocket.NewWithParams(p).Credentials(apiKey, apiSecret)
	err := c.Connect()
	if err != nil {
		log.Fatalf("connecting authenticated websocket: %s", err)
	}
	go func() {
		for msg := range c.Listen() {
			//bitfinex.WalletUpdate{}
			log.Printf("MSG RECV: %#v", msg)
		}
	}()

	//ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	//c.Authenticate(ctx)

	time.Sleep(time.Second * 10)
	return 0
}

func GetAvaliableBalance() float64{
	wallets, _ := bitfinexClient.Wallet.Wallet()
	return getFundingWalletAvalBalance(wallets)
}

func getFundingWalletAvalBalance(wallets *bitfinex.WalletSnapshot) float64{
	for _, wallet := range wallets.Snapshot{
		if wallet.Type == WalletType(Funding).String(){
			return wallet.BalanceAvailable
		}
	}
	log.Println("[Error] Didn't get funding wallet...")
	return -1
}
