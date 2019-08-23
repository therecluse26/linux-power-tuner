package conf

import (
	"encoding/json"
	"github.com/therecluse26/linux-power-tuner/pkg/utils"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

/*
 * Presets should be defined by pointing to events within "events" directory
 * within a .json preset file
 *
 * User-defined presets should be enabled/disabled by symlinking files
 * from "presets" into "enabled"
 */
type Config struct {
	HomePath string
	ConfigPath string
	GlobalConfigFile string
	ConfigEventsPath string
	AvailablePresetPath string
	EnabledPresetPath string
	CustomScriptsPath string
	AppConfig interface{} `json:"appConfig"`
}

func LoadGlobalConfig() *Config {
	userInfo, err := user.Current()
	if err != nil {
		utils.HandleError(err, 0, true, true)
	}
	c := new(Config)
	c.HomePath = userInfo.HomeDir + "/"
	c.ConfigPath = c.HomePath + ".config/power-tuner/"
	c.GlobalConfigFile = c.ConfigPath + "global-config.json"
	c.ConfigEventsPath = c.ConfigPath + "events/"
	c.AvailablePresetPath = c.ConfigPath + "available-presets/"
	c.EnabledPresetPath = c.ConfigPath + "enabled-presets/"
	c.CustomScriptsPath = c.ConfigPath + "custom-scripts/"

	return c
}

type Preset struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Events      []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PollingTime int64  `json:"pollingTime"`
		Reactions   []struct {
			Name     string `json:"name"`
			Function struct {
				Name string `json:"name"`
				Args []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"args"`
			} `json:"function"`
		} `json:"reactions"`
	} `json:"events"`
}

/**** Sample JSON Config file
{
	"name": "configName",
	"description": "configDesc",
	"events": [{
		"name": "eventName",
		"description": "eventDesc",
		"pollingTime": 2000000000000000,
		"reactions": [{
			"name":"reaction1",
			"function": {
				"name": "function1",
				"args": [
					{"key": "keyval1", "value": "valval1"},
					{"key": "keyval2", "value": "valval2"}
				]
			}
		}]
	}]
}
 */

/*
 * Loads active preset files from EnabledPresetPath directory
 */
func (c *Config) LoadActivePresets() []Preset {
	var presets []Preset
	_ = filepath.Walk(c.EnabledPresetPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			presets = append(presets, LoadPreset(path))
		}
		return nil
	})
	return presets
}

/*
 * Returns Preset struct from json file
 */
func LoadPreset(filePath string) Preset {
	presetFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		utils.HandleError(err, 0, true, true)
	}
	presetStruct := Preset{}
	err = json.Unmarshal(presetFile, &presetStruct)
	return presetStruct
}