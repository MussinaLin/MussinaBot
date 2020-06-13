package Bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"log"
	"time"
)

var AvailableBalance float64
var wsClient *websocket.Client

func GetAvailableBalance() float64{
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
	log.Println("Start Bitfinex WS...")
	go func() {
		for msg := range wsClient.Listen() {

			wu,ok := msg.(*bitfinex.WalletUpdate)
			if ok{
				log.Println("[Got bitfinex.WalletUpdate]")
				log.Printf("MSG RECV: %#v", wu)
				if wu.Type == WalletType(Funding).String(){
					AvailableBalance = wu.BalanceAvailable
				}
			}
		}
	}()

	//ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	//c.Authenticate(ctx)
}

func CloseWS(){
	log.Println("close websocket...")
	wsClient.Close()
	time.Sleep(3 * time.Second)
	log.Println("close websocket succ...")
}
