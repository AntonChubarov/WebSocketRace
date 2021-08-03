package app

import (
	"judge/domain"
	"log"
	"time"
)

type StepCommander struct {
	stepChannels *domain.StepChannelsStorage
	stepTicker   *time.Ticker
}

func NewStepCommander(stepChannels *domain.StepChannelsStorage) *StepCommander {
	return &StepCommander{stepChannels: stepChannels}
}

func (s *StepCommander) Run() {
	s.stepTicker = time.NewTicker(StepTime)
	var signal time.Time

	for {
		select {
		case signal = <- s.stepTicker.C:
			//log.Println("Step")
			for i := range s.stepChannels.Channels {
				s.stepChannels.Channels[i] <- signal
				log.Println("Sended to stepChannel", i)
			}
		default:
			continue
		}
		time.Sleep(LoopSleepTime)
	}
}
