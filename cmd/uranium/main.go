package main

import (
	"github.com/spf13/viper"
	"github.com/therecluse26/uranium/pkg/function"
	"github.com/therecluse26/uranium/pkg/utils"
	"os/user"
)

func main() {
	LoadConfig()

	gf := function.GenericFunction{}

	gf.FuncName("blahblahblah")

	/*
	availPresets := preset.GetAvailablePresets()
	var err error
	for _, p := range availPresets {
		err = preset.EnablePreset(p)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = preset.DisablePreset(availPresets[1])
	if err != nil {
		log.Fatal(err)
	}*/
	// Register events
}

func LoadConfig() {
	userInfo, err := user.Current()
	if err != nil {
		utils.HandleError(err, 0, true, true)
	}
	viper.SetDefault("HomePath", userInfo.HomeDir + "/")
	viper.SetDefault("ConfigPath",  viper.Get("HomePath").(string) + ".config/uranium/")
	viper.SetDefault("GlobalConfigFile", viper.Get("ConfigPath").(string) + "global-config.json")
	viper.SetDefault("ConfigEventsPath", viper.Get("ConfigPath").(string) + "events/")
	viper.SetDefault("AvailablePresetPath", viper.Get("ConfigPath").(string) + "available-presets/")
	viper.SetDefault("EnabledPresetPath", viper.Get("ConfigPath").(string) + "enabled-presets/")
	viper.SetDefault("CustomScriptsPath", viper.Get("ConfigPath").(string) + "custom-scripts/")
}

