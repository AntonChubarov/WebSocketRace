package domain

import "time"

type RacersInfoStorage struct {
	Info []RacerInfo
}

func NewRacersInfoStorage() *RacersInfoStorage {
	return &RacersInfoStorage{Info: []RacerInfo{}}
}

func (ris *RacersInfoStorage) AddRacer(info RacerInfo) {
	ris.Info = append(ris.Info, info)
}

type StepChannelsStorage struct {
	Channels []chan time.Time
}

func NewStepChannelsStorage() *StepChannelsStorage {
	return &StepChannelsStorage{Channels: []chan time.Time{}}
}

func (scs *StepChannelsStorage) AddChannel(channel chan time.Time) {
	scs.Channels = append(scs.Channels, channel)
}

type InfoChannelsStorage struct {
	Channels []chan RacerInfo
}

func NewInfoChannelsStorage() *InfoChannelsStorage {
	return &InfoChannelsStorage{Channels: []chan RacerInfo{}}
}

func (ics *InfoChannelsStorage) AddChannel(channel chan RacerInfo) {
	ics.Channels = append(ics.Channels, channel)
}

type DisplayChannelStorage struct {
	Channels []chan []RacerInfo
}

func NewDisplayChannelStorage() *DisplayChannelStorage {
	return &DisplayChannelStorage{Channels: []chan []RacerInfo{}}
}

func (dcs *DisplayChannelStorage) AddChannel(channel chan []RacerInfo) {
	dcs.Channels = append(dcs.Channels, channel)
}
