package Bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"log"
)

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
	return 0
}
