package infrastructure

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"racer/domain"
)

type WebSocketConnector struct {
	Connection *websocket.Conn
}

var addr = flag.String("addr", "localhost:8080", "http service address")

func NewWebSocketConnector() *WebSocketConnector {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	return &WebSocketConnector{
		Connection: c,
	}
}

func (w *WebSocketConnector) SendMessage(message domain.RacerInfo) {
	err := w.Connection.WriteJSON(message)
	if err != nil {
		log.Println("write:", err)
	}
}

func (w *WebSocketConnector) ReceiveMessage() domain.ServerInfo {
	var serverInfo domain.ServerInfo
	err := w.Connection.ReadJSON(&serverInfo)
	if err != nil {
		log.Println("read:", err)
	}
	return serverInfo
}
