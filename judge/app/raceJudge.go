package app

import (
	"fmt"
	"judge/domain"
	"log"
	"sort"
	"sync"
	"time"
)

type JudgeOfRace struct {
	RacersInfo      *domain.RacersInfoStorage
	StepChannels    *domain.StepChannelsStorage
	StepCommander    *StepCommander
	DisplayChannels  *domain.DisplayChannelStorage
	DisplayCommander *DisplayCommander
	InfoChannels    *domain.InfoChannelsStorage
	InfoCollector *InfoCollector
	StopChannel     chan bool
	IsInactiveRacer *[]bool
	InactiveCount   int
	MutexRacersInfo *sync.RWMutex
}

func NewRaceJudge(
	stopChannel chan bool,
	) *JudgeOfRace {
	ri := domain.NewRacersInfoStorage()
	stepChannel := domain.NewStepChannelsStorage()
	displayChannel := domain.NewDisplayChannelStorage()
	infoChannels := domain.NewInfoChannelsStorage()
	isInactive := &[]bool{}

	log.Println(fmt.Sprintf("%p", isInactive))

	mu := sync.RWMutex{}
	return &JudgeOfRace{
		RacersInfo:       ri,
		StepChannels:     stepChannel,
		StepCommander:    NewStepCommander(stepChannel),
		DisplayChannels:  displayChannel,
		DisplayCommander: NewDisplayCommander(ri, displayChannel, &mu),
		InfoChannels:     infoChannels,
		StopChannel:      stopChannel,
		IsInactiveRacer:  isInactive,
		MutexRacersInfo:  &mu,
		InfoCollector:    NewInfoCollector(ri, infoChannels, isInactive, &mu),
	}
}

func (j *JudgeOfRace) startRace() {
	go j.StepCommander.Run()
	go j.DisplayCommander.Run()
	go j.InfoCollector.Run()
	go j.startToJudge()
}

func (j *JudgeOfRace) startToJudge() {
	sortedInfo := make([]domain.RacerInfo, len(j.RacersInfo.Info))
	var nameOfRacerToStop string
	for {
		time.Sleep(LoopSleepTime)
		j.MutexRacersInfo.RLock()

		copy(sortedInfo, j.RacersInfo.Info)

		for i := range j.RacersInfo.Info {
			if !(*j.IsInactiveRacer)[i] && j.RacersInfo.Info[i].StepInLap >= 20 {
				(*j.IsInactiveRacer)[i] = true
				j.InactiveCount++
				fmt.Println(j.RacersInfo.Info[i].Name, "was too slow!")
				if j.InactiveCount == len(j.InfoChannels.Channels)-1 {
					break
				}
			}
		}

		j.MutexRacersInfo.RUnlock()
		sort.SliceStable(sortedInfo, func(i, j int) bool {
			return sortedInfo[i].Score > sortedInfo[j].Score
		})

		if j.InactiveCount < len(j.InfoChannels.Channels)-1 {
			if sortedInfo[len(sortedInfo)-1-j.InactiveCount].Lap < sortedInfo[len(sortedInfo)-2-j.InactiveCount].Lap {
				nameOfRacerToStop = sortedInfo[len(sortedInfo)-1-j.InactiveCount].Name
				racerIndex := j.findRacerIndexByName(nameOfRacerToStop)
				(*j.IsInactiveRacer)[racerIndex] = true
				j.InactiveCount++
			}
		}

		if j.InactiveCount == len(j.InfoChannels.Channels)-1 {
			j.MutexRacersInfo.RLock()
			fmt.Println("\nThe winner is " + sortedInfo[0].Name)
			for i := range sortedInfo{
				fmt.Println(sortedInfo[i].Name, "Score:", sortedInfo[i].Score)
			}
			j.MutexRacersInfo.RUnlock()
			j.StopChannel <- true
		}
	}
}

func (j *JudgeOfRace) findRacerIndexByName (name string) int {
	j.MutexRacersInfo.RLock()
	for i := range j.RacersInfo.Info {
		if j.RacersInfo.Info[i].Name == name {
			j.MutexRacersInfo.RUnlock()
			return i
		}
	}
	panic(fmt.Errorf("racer not found"))
}

func (j *JudgeOfRace) AddNewRacer(racer *RacerAgent)  {
	log.Println(fmt.Sprintf("%p", j.IsInactiveRacer))

	j.RacersInfo.AddRacer(racer.RacerInfo)
	stepChannel := make(chan time.Time)
	infoChannel := make(chan domain.RacerInfo)
	displayChannel := make(chan []domain.RacerInfo)

	j.StepChannels.AddChannel(stepChannel)
	j.InfoChannels.AddChannel(infoChannel)
	j.DisplayChannels.AddChannel(displayChannel)

	index := len(j.StepChannels.Channels)-1

	racer.SetChannels(&j.StepChannels.Channels[index], &j.InfoChannels.Channels[index], &j.DisplayChannels.Channels[index])

	racer.conn.WriteJSON(domain.ServerInfo{Message: NewCommand, StepsInLap: StepsInLap})

	go racer.StartRace()
	go racer.StartShowRaceSatus()

	log.Println(racer.RacerInfo.Name, "ready")

	if len(j.RacersInfo.Info) == Racers {
		*j.IsInactiveRacer = make([]bool, Racers)

		log.Println(fmt.Sprintf("%p", j.IsInactiveRacer))

		j.startRace()
	}
}
