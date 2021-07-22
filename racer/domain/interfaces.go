package domain

type Client interface {
	SendMessage(RacerInfo)
	ReceiveMessage() ServerInfo
}