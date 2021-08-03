package app

import (
	"judge/domain"
	"log"
	"sync"
	"time"
)

type DisplayCommander struct {
	racersInfo      *domain.RacersInfoStorage
	displayTicker   *time.Ticker
	displayChannels *domain.DisplayChannelStorage
	mutexRacersInfo *sync.RWMutex
}

func NewDisplayCommander(racersInfo *domain.RacersInfoStorage,
	displayChannel *domain.DisplayChannelStorage,
	mutexRacersInfo *sync.RWMutex,
) *DisplayCommander {
	return &DisplayCommander{racersInfo: racersInfo,
		displayChannels: displayChannel,
		mutexRacersInfo: mutexRacersInfo,
	}
}

func (d *DisplayCommander) Run() {
	d.displayTicker = time.NewTicker(DisplayTime)

	for {
		select {
		case <-d.displayTicker.C:
			log.Println("Display tick")
			d.mutexRacersInfo.RLock()
			for i := range d.displayChannels.Channels {
				d.displayChannels.Channels[i] <- d.racersInfo.Info
				log.Println("Info sent to display channel", i)
			}

			d.mutexRacersInfo.RUnlock()
		default:
			continue
		}
		time.Sleep(LoopSleepTime)
	}
}
