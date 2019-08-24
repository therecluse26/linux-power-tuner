package main

import (
	"github.com/spf13/viper"
	"os/user"
	"testing"
)

func TestLoadGlobalConfig (t *testing.T) {

	t.Run("check_global_config", func(t *testing.T) {
		userInfo, err := user.Current()
		if err != nil {
			t.Errorf("Error pulling current user")
		}
		homeFolder := userInfo.HomeDir + "/"

		LoadConfig()

		if viper.Get("HomePath").(string) != homeFolder {
			t.Errorf("got %s want %s", viper.Get("HomePath").(string), homeFolder)
		}
		confPath := homeFolder + ".config/uranium/"
		if viper.Get("ConfigPath").(string) != confPath {
			t.Errorf("got %s want %s", viper.Get("ConfigPath").(string), confPath)
		}
		globalConfFile := confPath + "global-config.json"
		if viper.Get("GlobalConfigFile").(string) != globalConfFile {
			t.Errorf("got %s want %s", viper.Get("GlobalConfigFile").(string), globalConfFile)
		}
		confEvtPath := confPath + "events/"
		if viper.Get("ConfigEventsPath").(string) != confEvtPath {
			t.Errorf("got %s want %s", viper.Get("ConfigEventsPath").(string), confEvtPath)
		}
		availPresets := confPath + "available-presets/"
		if viper.Get("AvailablePresetPath").(string) != availPresets {
			t.Errorf("got %s want %s", viper.Get("AvailablePresetPath").(string), availPresets)
		}
		enabPresets := confPath + "enabled-presets/"
		if viper.Get("EnabledPresetPath").(string) != enabPresets {
			t.Errorf("got %s want %s", viper.Get("EnabledPresetPath").(string), enabPresets)
		}
		custScripts := confPath + "custom-scripts/"
		if viper.Get("CustomScriptsPath").(string) != custScripts {
			t.Errorf("got %s want %s", viper.Get("CustomScriptsPath").(string), custScripts)
		}
	})


}