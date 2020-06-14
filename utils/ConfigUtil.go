package utils

import (
	"MussinaBot/encryption"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)



type Config struct {
	ApiKey string
	ApiSecret string
	PubEndpoint string
	MinLoan float64
	BalanceLeft float64
	FrrBias float64
	FrrLoanMonthRate float64
	FrrIncreaseRate float64
	FrrCalculatePriorSecs int
	MaxSingleOrderAmount float64
	OrdersNotLendTh int
	WsURL string
}

func LoadConfig(envPath... string) (*Config, error){
	// get private key from system env
	key := os.Getenv("key")
	if len(key) > 0 && key[0] == 'M' {
		log.Println("get private key succ...")
	}else{
		return nil, fmt.Errorf("get private key fail")
	}

	// load .env file
	var err error = nil
	if len(envPath) > 0 { //for test case read .env file
		err = godotenv.Load(envPath[0])
	}else{
		err = godotenv.Load()
	}
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	// platform
	encryApiKey := os.Getenv("apiKey")
	encryApiSecret := os.Getenv("apiSecret")
	apiKey := encryption.Decrypt(key, encryApiKey)
	apiSecret := encryption.Decrypt(key, encryApiSecret)
	pubEndpoint := os.Getenv("bf.pub.endpoint")

	// balance
	minLoad := os.Getenv("bf.balance.min.loan")
	balanceLeft := os.Getenv("bf.balance.left")

	//FRR
	frrBias := os.Getenv("bf.FRR.bias")
	frrLoanMonthRate := os.Getenv("bf.FRR.loan.month")
	frrIncreaseRate := os.Getenv("bf.FRR.increasing.rate")
	frrCalculatePriorSecs := os.Getenv("bf.FRR.calculate.prior.seconds")

	// order
	maxSingleOrderAmount := os.Getenv("bf.order.single.max.amount")
	ordersNotLendTh := os.Getenv("bf.order.not.lend.threshold")
	//WS
	wsUrl := os.Getenv("bf.ws.url")

	cfg := &Config{}
	cfg.ApiKey = apiKey
	cfg.ApiSecret = apiSecret
	cfg.PubEndpoint = pubEndpoint
	cfg.MinLoan, _ = strconv.ParseFloat(minLoad, 64)
	cfg.FrrBias, _ = strconv.ParseFloat(frrBias, 64)
	cfg.FrrLoanMonthRate, _ = strconv.ParseFloat(frrLoanMonthRate, 64)
	cfg.FrrIncreaseRate, _ = strconv.ParseFloat(frrIncreaseRate, 64)
	cfg.FrrCalculatePriorSecs,_ = strconv.Atoi(frrCalculatePriorSecs)
	cfg.BalanceLeft, _ = strconv.ParseFloat(balanceLeft, 64)
	cfg.MaxSingleOrderAmount, _ = strconv.ParseFloat(maxSingleOrderAmount, 64)
	cfg.WsURL = wsUrl
	cfg.OrdersNotLendTh, _ = strconv.Atoi(ordersNotLendTh)
	log.Println("get config succ...")
	return cfg,nil
}
