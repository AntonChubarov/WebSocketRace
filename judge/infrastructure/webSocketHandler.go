package infrastructure

import (
	"github.com/gorilla/websocket"
	"judge/app"
	"judge/domain"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{} // use default options

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var c *websocket.Conn
	c, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	var inMessage domain.RacerInfo

	for {
		err = c.ReadJSON(&inMessage)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv from: %s", inMessage.Name)

		app.RacerInfoChannel <- inMessage

		outMessage := <- app.ServerInfoChannel

		err = c.WriteJSON(outMessage)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}


}