package main

import (
	"flag"
	"judge/app"
	"judge/domain"
	"judge/infrastructure"
	"log"
	"net/http"
)

//var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	var racerInfoChannel = make(chan domain.RacerInfo)
	var serverInfoChannel = make(chan domain.ServerInfo)

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", infrastructure.Handle)

	go app.StartMessageProcessing()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
