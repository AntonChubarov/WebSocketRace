package app

import (
	"fmt"
	"judge/domain"
	"log"
	"sync"
	"time"
)

type InfoCollector struct {
	racersInfo      *domain.RacersInfoStorage
	infoChannels    *domain.InfoChannelsStorage
	isInactiveRacer *[]bool
	mutexRacersInfo *sync.RWMutex
}

func NewInfoCollector(racersInfo *domain.RacersInfoStorage,
	infoChannels *domain.InfoChannelsStorage,
	isInactiveRacer *[]bool,
	mutexRacersInfo *sync.RWMutex,
) *InfoCollector {

	log.Println(fmt.Sprintf("%p", isInactiveRacer))

	return &InfoCollector{racersInfo: racersInfo,
		infoChannels: infoChannels,
		isInactiveRacer: isInactiveRacer,
		mutexRacersInfo: mutexRacersInfo,
	}
}


func (ic *InfoCollector) Run() {
	var ok bool
	var in domain.RacerInfo
	for {
		for i := range ic.infoChannels.Channels {
			if in, ok = <- ic.infoChannels.Channels[i]; ok && !(*ic.isInactiveRacer)[i] {
				ic.mutexRacersInfo.Lock()
				ic.racersInfo.Info[i] = in
				ic.mutexRacersInfo.Unlock()
			}
		}
		time.Sleep(LoopSleepTime)
	}
}