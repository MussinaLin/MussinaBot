package http

import (
	"MussinaBot/api"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	http_server *http.Server
}


func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func NewHttpServer() *HttpServer{
	log.Println("[NewHttpServer...]")
	return &HttpServer{}
}

func StartHttpServer(){
	log.Println("[StartHttpServer...]")
	//httpServer := &http.Server{
	//	Addr:    addr,
	//	Handler: nil,
	//}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/balance", api.GetGeneralResp)
	http.HandleFunc("/funds/provided", api.GetProvideFunds)
	http.HandleFunc("/funds/earned", api.GetEarnedInterest)
	http.HandleFunc("/config", api.ReloadConfig)
	go func() {
		http.ListenAndServe(":8090", nil)
	}()
}
