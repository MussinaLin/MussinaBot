package Bitfinex

import (
	"MussinaBot/utils"
	"testing"
)

func TestGetAvaliableBalance(t *testing.T){
	cfg, _ := utils.LoadConfig("/Users/Mussina/GitRepo/MussinaBot/.env")
	SetConfig(cfg.ApiKey, cfg.ApiSecret, "https://api.bitfinex.com/v2/")
	avalBalance := GetAvaliableBalance()
	t.Log("avaliable balance:", avalBalance)
}


//func TestGetAvaliableBalance(t *testing.T) {
//	tests := []struct {
//		name string
//		want float64
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := GetAvaliableBalance(); got != tt.want {
//				t.Errorf("GetAvaliableBalance() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_getFundingWalletAvalBalance(t *testing.T) {
//	type args struct {
//		wallets *bitfinex.WalletSnapshot
//	}
//	tests := []struct {
//		name string
//		args args
//		want float64
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := getFundingWalletAvalBalance(tt.args.wallets); got != tt.want {
//				t.Errorf("getFundingWalletAvalBalance() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}