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
	viper.SetDefault("ProjectLogo", viper.Get("ConfigPath").(string) + "react-logo.png")
	viper.SetDefault("GlobalConfigFile", viper.Get("ConfigPath").(string) + "global-config.json")
	viper.SetDefault("ConfigEventsPath", viper.Get("ConfigPath").(string) + "events/")
	viper.SetDefault("AvailablePresetPath", viper.Get("ConfigPath").(string) + "available-presets/")
	viper.SetDefault("EnabledPresetPath", viper.Get("ConfigPath").(string) + "enabled-presets/")
	viper.SetDefault("CustomScriptsPath", viper.Get("ConfigPath").(string) + "custom-scripts/")
}

