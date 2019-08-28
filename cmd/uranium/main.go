package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/therecluse26/uranium/pkg/function"
	"log"
	"os/user"
	"path/filepath"
	"runtime"
)

func main() {
	LoadConfig()

	function.LoadFunctions()

	ret, err := function.CallFunction(function.UserFuncs, "Multiply",  1, 3, 4)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)

	uc, err := function.CallFunction(function.UserFuncs, "UpperCase", "BlaHHblah")
	fmt.Println(uc)


	/*_, e := function.CallFunction(function.EventFuncs,"UpperCase", "testing, 123")
	if e != nil {
		panic(e)
	}
	//fmt.Println(v)

	_, e = function.CallFunction(function.ReactionFuncs,"Multiply", 2, 5, 6)
	if e != nil {
		panic(e)
	}*/

	//fmt.Println(r)

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
		log.Fatal(err)
	}
	_, b, _, _ := runtime.Caller(0)

	basePath := filepath.Dir(filepath.Dir(filepath.Dir(b)))
	viper.SetDefault("HomePath", userInfo.HomeDir + "/")
	viper.SetDefault("ProjectBase", basePath)
	viper.SetDefault("ConfigPath",  viper.Get("HomePath").(string) + ".config/uranium/")
	viper.SetDefault("GlobalConfigFile", viper.Get("ConfigPath").(string) + "global-config.json")
	viper.SetDefault("ConfigEventsPath", viper.Get("ConfigPath").(string) + "events/")
	viper.SetDefault("AvailablePresetPath", viper.Get("ConfigPath").(string) + "available-presets/")
	viper.SetDefault("EnabledPresetPath", viper.Get("ConfigPath").(string) + "enabled-presets/")
	viper.SetDefault("CustomScriptsPath", viper.Get("ConfigPath").(string) + "custom-scripts/")
}

