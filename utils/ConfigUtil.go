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

	encryApiKey := os.Getenv("apiKey")
	encryApiSecret := os.Getenv("apiSecret")
	apiKey := encryption.Decrypt(key, encryApiKey)
	apiSecret := encryption.Decrypt(key, encryApiSecret)
	pubEndpoint := os.Getenv("bf.pub.endpoint")
	minLoad := os.Getenv("bf.min.loan")

	cfg := &Config{}
	cfg.ApiKey = apiKey
	cfg.ApiSecret = apiSecret
	cfg.PubEndpoint = pubEndpoint
	cfg.MinLoan, _ = strconv.ParseFloat(minLoad, 64)

	log.Println("get config succ...")
	return cfg,nil
}
