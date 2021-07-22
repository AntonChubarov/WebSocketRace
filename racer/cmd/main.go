package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"racer/app"
	"racer/infrastructure"
	"time"
)



func main() {
	flag.Parse()
	log.SetFlags(0)

	fmt.Println("Enter racer's name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	racer := app.NewRandomStepRacer(name)

	client := infrastructure.NewWebSocketConnector()

	console := infrastructure.NewConsole()

	handler := app.NewMessageHandler(racer, client, console)

	go handler.StartHandling()

	for{
		time.Sleep(time.Second)
	}
}
