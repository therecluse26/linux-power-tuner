package main

import (
	"github.com/therecluse26/linux-power-tuner/pkg/conf"
	"log"
)

func main() {

	config := conf.LoadConfig()
	availPresets := config.GetAvailablePresets()
	var err error
	for _, p := range availPresets {
		err = config.EnablePreset(p)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = config.DisablePreset(availPresets[1])
	if err != nil {
		log.Fatal(err)
	}

	// Register events

}
