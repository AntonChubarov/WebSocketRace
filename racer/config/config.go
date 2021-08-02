package config

var stepsInLap int

func SetStepsPerLapValue(value int) {
	stepsInLap = value
}

func StepsInLap() int {
	return stepsInLap
}
