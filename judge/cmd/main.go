package main

import (
	"flag"
	"judge/app"
	"judge/infrastructure"
	"log"
	"net/http"
)

//var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", infrastructure.WebSocketHandler)

	go app.StartMessageHandling()

	log.Fatal(http.ListenAndServe(":8080", nil))


}