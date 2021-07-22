package app

import "judge/domain"

var RacerInfoChannel = make(chan domain.RacerInfo)
var ServerInfoChannel = make(chan domain.ServerInfo)

func StartMessageHandling() domain.ServerInfo {
	for {
		inMessage := <-RacerInfoChannel

		switch inMessage.Message {
		case NewCommand:
			ServerInfoChannel <- domain.ServerInfo{Message: NewCommand}
		case UpdateCommand:

		default:

		}
	}
}
