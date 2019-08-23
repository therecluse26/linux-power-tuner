package main

import (
	"fmt"
	"github.com/therecluse26/linux-power-tuner/pkg/conf"
)

func main() {

	config := conf.LoadGlobalConfig()
	presets := config.LoadActivePresets()
	for _, p := range presets {
		fmt.Println("Preset: " + p.Name + " - " + p.Description)
		for _, e := range p.Events {
			fmt.Println("Event: " + e.Name + " - " + e.Description)
			fmt.Println(e.PollingTime)
			for _, r := range e.Reactions {
				fmt.Println("Reaction: " + r.Name)
				fmt.Println("Function: " + r.Function.Name)
				for _, a := range r.Function.Args {
					fmt.Println(a.Key + " - " + a.Value)
				}
			}
		}
	}

	// Register events

}
