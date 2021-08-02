package app

import "judge/domain"

type MessageHandler struct {
	RacerInfoChannel chan domain.RacerInfo
	ServerInfoChannel chan domain.ServerInfo
	Judge JudgeOfRace
}

func NewMessageHandler(racerInfoChannel chan domain.RacerInfo,
	serverInfoChannel chan domain.ServerInfo,
	judge JudgeOfRace,
	) *MessageHandler {
	return &MessageHandler{RacerInfoChannel: racerInfoChannel,
		ServerInfoChannel: serverInfoChannel,
		Judge: judge,
	}
}

func(m *MessageHandler) StartMessageProcessing() domain.ServerInfo {
	for {
		inMessage := <- m.RacerInfoChannel

		switch inMessage.Message {
		case NewCommand:
			m.ServerInfoChannel <- domain.ServerInfo{Message: NewCommand}
		case UpdateCommand:

		default:

		}
	}
}
