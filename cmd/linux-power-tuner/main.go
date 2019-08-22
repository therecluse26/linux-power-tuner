package main

import (
	"fmt"
	"github.com/therecluse26/linux-power-tuner/pkg/system"
)

func main() {
	//sysInfo := system.GetInfo()
	//fmt.Println(sysInfo.Cpu.EnabledCores)

	batts := system.GetBatteries()
	fmt.Println(batts.GetBatteryPercent())
	fmt.Println(batts.GetChargingRate())

	// Register events

}
