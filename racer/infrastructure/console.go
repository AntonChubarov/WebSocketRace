package infrastructure

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"racer/domain"
	"runtime"
	"strconv"
)

type Console struct {

}

func NewConsole() *Console {
	return &Console{
	}
}

func (c *Console) ShowRaceInfo(info domain.ServerInfo)  {
	CallClear()
	racersInfo := info.Info
	for i := range racersInfo {
		printRacerString(racersInfo[i], info.StepsInLap)
	}
}

func printRacerString(racerInfo domain.RacerInfo, stepsInLap int) {
	racerString := ""
	racerString += racerInfo.Name
	racerString += " |"
	for i := 0; i < stepsInLap; i++ {
		if i < racerInfo.Score % stepsInLap {
			racerString += "-"
		}
		if i == racerInfo.Score % stepsInLap {
			racerString += ">"
		}
		if i > racerInfo.Score % stepsInLap {
			racerString += " "
		}
	}
	racerString += " | "
	racerString += "Lap: "
	racerString += strconv.Itoa(1 + racerInfo.Score / stepsInLap)
	fmt.Println(racerString)
}

func (c *Console) ShowMessage(message string) {
	fmt.Println(message)
}

// Console clear

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println("unable to perform console command")
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println("unable to perform console command")
		}
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else { // unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}