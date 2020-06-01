package main

import (
	"log"
	"os"
)

func main() {
	key := os.Getenv("key")
	log.Println("get key:", key)
	log.Println("Starting the application...")
	//plaintext := encryption.Decrypt(key, ciphertext)
	//fmt.Printf("Decrypted: %s\n", plaintext)

	//key := "BmkGCnViEQ_FsFlqdk6li74oDVH41fag7dwdVQt8CIL9FlDst2ebmh_Zm3eXRSN-YxPpf6K4uOLZm74="
	//secret := "Zhz0MxMkxPQ0VfBai-Wzi6nxPGaCw3bYYxizHzgpFnMpY5Hoa2F0WGWQv66nl6ayrHpSUbO9nIgeISw="
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
