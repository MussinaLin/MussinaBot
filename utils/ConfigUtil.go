package utils

import (
	"MussinaBot/encryption"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)



type Config struct {
	ApiKey string
	ApiSecret string
	PubEndpoint string
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

	cfg := &Config{}
	cfg.ApiKey = apiKey
	cfg.ApiSecret = apiSecret
	cfg.PubEndpoint = pubEndpoint
	return cfg,nil
}
