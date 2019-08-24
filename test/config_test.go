package main

import (
	"github.com/therecluse26/linux-power-tuner/pkg/conf"
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
		config := conf.LoadConfig()
		if config.HomePath != homeFolder {
			t.Errorf("got %s want %s", config.HomePath, homeFolder)
		}
		confPath := homeFolder + ".config/power-tuner/"
		if config.ConfigPath != confPath {
			t.Errorf("got %s want %s", config.ConfigPath, confPath)
		}
		globalConfFile := confPath + "global-config.json"
		if config.GlobalConfigFile != globalConfFile {
			t.Errorf("got %s want %s", config.GlobalConfigFile, globalConfFile)
		}
		confEvtPath := confPath + "events/"
		if config.ConfigEventsPath != confEvtPath {
			t.Errorf("got %s want %s", config.ConfigEventsPath, confEvtPath)
		}
		availPresets := confPath + "available-presets/"
		if config.AvailablePresetPath != availPresets {
			t.Errorf("got %s want %s", config.AvailablePresetPath, availPresets)
		}
		enabPresets := confPath + "enabled-presets/"
		if config.EnabledPresetPath != enabPresets {
			t.Errorf("got %s want %s", config.EnabledPresetPath, enabPresets)
		}
		custScripts := confPath + "custom-scripts/"
		if config.CustomScriptsPath != custScripts {
			t.Errorf("got %s want %s", config.CustomScriptsPath, custScripts)
		}
	})


}