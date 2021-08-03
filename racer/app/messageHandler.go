package app

import (
	"log"
	"racer/config"
	"racer/domain"
	"racer/infrastructure"
	"strconv"
)

type MessageHandler struct {
	Racer *RandomStepRacer
	Client domain.Client
	Console *infrastructure.Console
}

func NewMessageHandler(racer *RandomStepRacer, client domain.Client, console *infrastructure.Console) *MessageHandler {
	handler := &MessageHandler{
		Racer: racer,
		Client: client,
		Console: console,
	}

	message := handler.Racer.GetInfo()
	message.Message = NewCommand

	handler.Client.SendMessage(message)

	return handler
}

func (m *MessageHandler) StartHandling() {

	for {
		inMessage := m.Client.ReceiveMessage()

		switch inMessage.Message {
		case NewCommand:
			config.SetStepsPerLapValue(inMessage.StepsInLap)
			m.Console.ShowMessage("You are connected!\nSteps in lap = " + strconv.Itoa(inMessage.StepsInLap))
		case StepCommand:
			m.Racer.MakeStep()
			log.Println("stepMessage")
			outMessage := m.Racer.GetInfo()
			m.Client.SendMessage(outMessage)
		case UpdateCommand:
			log.Println("displayMessage")
			m.Console.ShowRaceInfo(inMessage)
		default:
			continue
		}
	}
}


