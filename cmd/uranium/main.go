package main

import (
	"github.com/spf13/viper"
	"github.com/therecluse26/uranium/pkg/function"
	"github.com/therecluse26/uranium/pkg/preset"
	"log"
	"os/user"
	"path/filepath"
	"runtime"
)

func main() {

	LoadConfig()
	presets := preset.LoadActivePresets()
	allFunctions := function.LoadFunctions()

	preset.RunPresets(presets, allFunctions)

	/*
	uc, _ := function.CallFunction(function.BuiltInFuncs, "Hello")
	fmt.Println(uc)

	t, _ := function.CallFunction(function.UserFuncs, "UpperCase",  "test. yeah yeah no. yeah")
	fmt.Println(t)

	ret, _ := function.CallFunction(function.BuiltInFuncs, "AddInts",  []int{1, 3, 4})
	fmt.Println(ret)

	m, _ := function.CallFunction(function.UserFuncs, "Multiply",  []int{1, 3, 4})
	fmt.Println(m)
	*/

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

