package domain

type RacerInfo struct {
	Message string `json:"message"`
	Name string `json:"name"`
	Step int `json:"step"`
	StepInLap int `json:"step_in_lap"`
	Score int `json:"score"`
	Lap int `json:"lap"`
}

type ServerInfo struct {
	Message string `json:"message"`
	StepsInLap int `json:"steps_in_lap"`
	Info []RacerInfo `json:"info"`
}