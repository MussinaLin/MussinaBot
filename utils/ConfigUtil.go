package utils

import (
	"MussinaBot/encryption"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var MussinaBotCfg *Config = &Config{}

type Config struct {
	ApiKey string `json:"-"`
	ApiSecret string `json:"-"`
	PubEndpoint string `json:"pub_endpoint"`
	MinLoan float64 `json:"min_loan"`
	BalanceLeft float64 `json:"balance_left"`
	FrrBias float64 `json:"frr_bias"`
	FrrLoanMonthRate float64 `json:"frr_loan_month_rate"`
	FrrIncreaseRate float64 `json:"frr_increase_rate"`
	FrrCalculatePriorSecs int `json:"frr_calculate_prior_secs"`
	MaxSingleOrderAmount float64 `json:"max_single_order_amount"`
	OrdersNotLendTh int `json:"orders_not_lend_th"`
	WsURL string `json:"ws_url"`
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
		err = godotenv.Overload(envPath[0])
	}else{
		err = godotenv.Overload()
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

	MussinaBotCfg.ApiKey = apiKey
	MussinaBotCfg.ApiSecret = apiSecret
	MussinaBotCfg.PubEndpoint = pubEndpoint
	MussinaBotCfg.MinLoan, _ = strconv.ParseFloat(minLoad, 64)
	MussinaBotCfg.FrrBias, _ = strconv.ParseFloat(frrBias, 64)
	MussinaBotCfg.FrrLoanMonthRate, _ = strconv.ParseFloat(frrLoanMonthRate, 64)
	MussinaBotCfg.FrrIncreaseRate, _ = strconv.ParseFloat(frrIncreaseRate, 64)
	MussinaBotCfg.FrrCalculatePriorSecs,_ = strconv.Atoi(frrCalculatePriorSecs)
	MussinaBotCfg.BalanceLeft, _ = strconv.ParseFloat(balanceLeft, 64)
	MussinaBotCfg.MaxSingleOrderAmount, _ = strconv.ParseFloat(maxSingleOrderAmount, 64)
	MussinaBotCfg.WsURL = wsUrl
	MussinaBotCfg.OrdersNotLendTh, _ = strconv.Atoi(ordersNotLendTh)
	printCfg(MussinaBotCfg)
	return MussinaBotCfg,nil
}

func printCfg(cfg *Config){
	log.Println("========== MussinaBot Config ==========")
	log.Println(fmt.Sprintf("PubEndpoint=%s", cfg.PubEndpoint))
	log.Println(fmt.Sprintf("MinLoan=%f", cfg.MinLoan))
	log.Println(fmt.Sprintf("FrrBias=%f", cfg.FrrBias))
	log.Println(fmt.Sprintf("FrrLoanMonthRate=%f", cfg.FrrLoanMonthRate))
	log.Println(fmt.Sprintf("FrrIncreaseRate=%f", cfg.FrrIncreaseRate))
	log.Println(fmt.Sprintf("FrrCalculatePriorSecs=%d", cfg.FrrCalculatePriorSecs))
	log.Println(fmt.Sprintf("BalanceLeft=%f", cfg.BalanceLeft))
	log.Println(fmt.Sprintf("MaxSingleOrderAmount=%f", cfg.MaxSingleOrderAmount))
	log.Println(fmt.Sprintf("WsURL=%s", cfg.WsURL))
	log.Println(fmt.Sprintf("OrdersNotLendTh=%d", cfg.OrdersNotLendTh))
	log.Println("=========================================")
}
