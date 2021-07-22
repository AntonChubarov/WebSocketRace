package domain

type RacerInfo struct {
	Message string `json:"message"`
	Name string `json:"name"`
	Step int `json:"step"`
	Score int `json:"score"`
}

type ServerInfo struct {
	Message string `json:"message"`
	StepsInLap int `json:"steps_in_lap"`
	Info []RacerInfo `json:"info"`
}