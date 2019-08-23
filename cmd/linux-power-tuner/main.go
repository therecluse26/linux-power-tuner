package main

import (
	"github.com/therecluse26/linux-power-tuner/pkg/conf"
	"log"
)

func main() {

	config := conf.LoadGlobalConfig()

	availPresets := config.GetAvailablePresets()

	err := config.EnablePreset(availPresets[1])
	if err != nil {
		log.Fatal(err)
	}
	err = config.DisablePreset(availPresets[1])
	if err != nil {
		log.Fatal(err)
	}

	// Register events

}
