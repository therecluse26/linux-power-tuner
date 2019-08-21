package main

import (
	"fmt"
	"github.com/therecluse26/linux-power-tuner/pkg/events/conditions"
)

func main() {
	//sysInfo := system.GetInfo()
	//fmt.Println(sysInfo.Cpu.EnabledCores)

	file := conditions.FileMeta{Path: "/home/brad/Downloads/search_test.txt"}

	fmt.Println(file.SearchFileValue(conditions.Search{Type: conditions.Simple, Query: 1}))


	// Register events

}
