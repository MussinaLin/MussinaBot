package main

import (
	"github.com/joho/godotenv"
	"MussinaBot/encryption"
	"log"
	"os"
)

func main() {
	key := os.Getenv("key")
	log.Println("get key:", key)
	log.Println("Starting the application...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	encryApiKey := os.Getenv("apiKey")
	encryApiSecret := os.Getenv("apiSecret")
	apiKey := encryption.Decrypt(key, encryApiKey)
	apiKey = apiKey
	apiSecret := encryption.Decrypt(key, encryApiSecret)
	apiSecret = apiSecret

	//c := rest.NewClientWithURL("https://api.bitfinex.com/v2/").Credentials(key, secret)
	//
	//snapHist, err := c.Funding.OfferHistory("fUSD")
	//if err != nil {
	//	panic(err)
	//}
	//for _, item := range snapHist.Snapshot {
	//	fmt.Println(item)
	//}
}
