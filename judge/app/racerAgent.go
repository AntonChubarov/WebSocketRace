package app

import (
	"github.com/gorilla/websocket"
	"judge/domain"
	"log"
	"time"
)

type RacerAgent struct {
	RacerInfo domain.RacerInfo
	conn     *websocket.Conn
	stepChan *chan time.Time
	infoChan       *chan domain.RacerInfo
	displayChannel *chan []domain.RacerInfo
}

func NewRacerAgent(conn *websocket.Conn) *RacerAgent {
	var inMessage domain.RacerInfo
	err := conn.ReadJSON(&inMessage)
	if err != nil {
		log.Println("read:", err)
	}
	return &RacerAgent{
		RacerInfo: inMessage,
		conn: conn,
	}
}

func (r *RacerAgent) StartRace() {
	var err error
	var inMessage domain.RacerInfo

Loop:
	for {
		<- *r.stepChan

		log.Println(r.RacerInfo.Name, "has received from step chan")

		outMessage := domain.ServerInfo{Message: StepCommand, StepsInLap: StepsInLap}

		err = r.conn.WriteJSON(outMessage)
		if err != nil {
			log.Println("write:", err)
			break Loop
		}

		err = r.conn.ReadJSON(&inMessage)
		if err != nil {
			log.Println("read:", err)
			break Loop
		}
		log.Printf("received from: %s", inMessage.Name)

		*r.infoChan <- inMessage
		time.Sleep(LoopSleepTime)
	}
}

func (r *RacerAgent) StartShowRaceSatus() {
	var err error
	var info []domain.RacerInfo
Loop:
	for {
		info = <- *r.displayChannel

		log.Println(r.RacerInfo.Name, "has receive from display channel")

		outMessage := domain.ServerInfo{Message: UpdateCommand, StepsInLap: StepsInLap, Info: info}

		err = r.conn.WriteJSON(outMessage)
		if err != nil {
			log.Println("write:", err)
			break Loop
		}

		log.Println("Display info sent to", r.RacerInfo.Name)

		time.Sleep(LoopSleepTime)
	}
}

func (r *RacerAgent) SetChannels(stepChan *chan time.Time,
	infoChan *chan domain.RacerInfo,
	displayChannel *chan []domain.RacerInfo,
	) {
	r.stepChan = stepChan
	r.infoChan = infoChan
	r.displayChannel = displayChannel
}
