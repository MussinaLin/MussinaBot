package Bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"log"
)

var AvailableBalance float64
var wsClient *websocket.Client

func getAvailableBalance() float64{
	return AvailableBalance
}

func StartAvaliableBalanceWS(apiKey string, apiSecret string, uri string){
	p := websocket.NewDefaultParameters()

	//p.URL = uri
	wsClient = websocket.NewWithParams(p).Credentials(apiKey, apiSecret)
	err := wsClient.Connect()
	if err != nil {
		log.Fatalf("connecting authenticated websocket: %s", err)
	}
	go func() {
		for msg := range wsClient.Listen() {
			log.Printf("MSG RECV: %#v", msg)
			_,ok := msg.(*bitfinex.WalletUpdate)
			if ok{
				log.Println("Got &bitfinex.WalletUpdate!")
			}
		}
	}()

	//ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	//c.Authenticate(ctx)
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
