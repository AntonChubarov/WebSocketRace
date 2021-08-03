package infrastructure

import (
	"github.com/gorilla/websocket"
	"judge/app"
	"log"
	"net/http"
	"time"
)

type WebSocketHandler struct {
	//RacerInfoChannel chan domain.RacerInfo
	//ServerInfoChannel chan domain.ServerInfo
	Judge *app.JudgeOfRace
}

func NewWebSocketHandler(judge *app.JudgeOfRace) *WebSocketHandler {
	return &WebSocketHandler{Judge: judge}
}

var upgrader = websocket.Upgrader{} // use default options

func (wsh *WebSocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var err error
	var c *websocket.Conn
	c, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	racer := app.NewRacerAgent(c)

	wsh.Judge.AddNewRacer(racer)

	for {
		time.Sleep(app.LoopSleepTime)
	}

	//var inMessage domain.RacerInfo
	//
	//for {
	//	err = c.ReadJSON(&inMessage)
	//	if err != nil {
	//		log.Println("read:", err)
	//		break
	//	}
	//	log.Printf("recv from: %s", inMessage.Name)
	//
	//	//select {
	//	//case app.RacerInfoChannel <- inMessage:
	//	//	continue
	//	//default:
	//	//	continue
	//	//}
	//
	//	RacerInfoChannel <- inMessage
	//
	//	//var outMessage domain.ServerInfo
	//
	//	//select {
	//	//case outMessage =  <- app.ServerInfoChannel:
	//	//	err = c.WriteJSON(outMessage)
	//	//	if err != nil {
	//	//		log.Println("write:", err)
	//	//		break
	//	//	}
	//	//default:
	//	//	continue
	//	//}
	//
	//	outMessage := <- app.ServerInfoChannel
	//
	//	err = c.WriteJSON(outMessage)
	//	if err != nil {
	//		log.Println("write:", err)
	//		break
	//	}
	//	time.Sleep(app.LoopSleepTime)
	//}
}