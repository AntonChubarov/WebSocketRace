package app

import (
	"racer/config"
	"racer/domain"
	"racer/infrastructure"
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
			m.Console.ShowMessage("You are connected")
		case StepCommand:
			m.Racer.MakeStep()
			outMessage := m.Racer.GetInfo()
			m.Client.SendMessage(outMessage)
		case UpdateCommand:
			m.Console.ShowRaceInfo(inMessage)
		default:
			continue
		}
	}
}


