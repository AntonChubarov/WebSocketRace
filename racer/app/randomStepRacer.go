package app

import (
	"crypto/rand"
	"math/big"
	"racer/config"
	"racer/domain"
)

type RandomStepRacer struct {
	Name     string
	Step     int
	StepInLap int
	Score    int
	Lap      int
}

func NewRandomStepRacer(name string) *RandomStepRacer {
	return &RandomStepRacer{
		Name:     name,
		Step:     0,
		StepInLap: 0,
		Score:    0,
		Lap: 1,
	}
}

func (r *RandomStepRacer) MakeStep() {
	points := 1 + randomInt(4)
	r.Step++
	r.Score += points
	prevLapValue := r.Lap
	r.Lap = 1 + r.Score / config.StepsInLap()
	if prevLapValue == r.Lap {
		r.StepInLap++
	} else {
		r.StepInLap = 0
	}
}

func (r *RandomStepRacer) GetInfo() domain.RacerInfo {
	return domain.RacerInfo{
		Message: UpdateCommand,
		Name: r.Name,
		Step: r.Step,
		Score: r.Score,
	}
}

func randomInt(max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return int(n)
}