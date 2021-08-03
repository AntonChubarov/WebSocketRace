package main

import (
	"flag"
	"fmt"
	"judge/app"
	"judge/infrastructure"
	"log"
	"net/http"
)

//var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	//var racerInfoChannel = make(chan domain.RacerInfo)
	//var serverInfoChannel = make(chan domain.ServerInfo)

	stopChannel := make (chan bool)

	judge := app.NewRaceJudge(stopChannel)
	log.Println(fmt.Sprintf("%p", judge.IsInactiveRacer))

	webSocketHandler := infrastructure.NewWebSocketHandler(judge)

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", webSocketHandler.Handle)

	//go app.StartMessageProcessing()

	log.Fatal(http.ListenAndServe(":8080", nil))

	for {
		<- stopChannel
		return
	}
}
