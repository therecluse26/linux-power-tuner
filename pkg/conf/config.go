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

func LoadConfig() *Config {
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

type Function struct {
	Name string `json:"name"`
	Args []struct {
		Key   string `json:"key"`
		Value interface{} `json:"value"`
	} `json:"args"`
}

type Preset struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Events      []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PollingTime int64  `json:"pollingTime"`
		Conditions []struct {
			Id 			int    		`json:"id"`
			Description string 		`json:"description"`
			Function 	Function	`json:"function"`
			ExpectedVal interface{} `json:"expected_val"`
		} `json:"conditions"`
		ConditionExp string `json:"condition_exp"`
	} `json:"events"`
	Reactions   []struct {
		Name     string		`json:"name"`
		Function Function	`json:"function"`
	} `json:"reactions"`
}

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

/*
 * Gets available preset file names from AvailablePresetPath
 */
func (c *Config) GetAvailablePresets() []string {
	var presetFiles []string
	_ = filepath.Walk(c.AvailablePresetPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			presetFiles = append(presetFiles, info.Name())
		}
		return nil
	})
	return presetFiles
}

/*
 * Enables preset by creating a link from the AvailablePresetPath to the EnabledPresetPath
 */
func (c *Config) EnablePreset(fileName string) error {
	err := os.Link(c.AvailablePresetPath + fileName, c.EnabledPresetPath + fileName)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return nil
}

/*
 * Disables preset by removing link file from EnabledPresetPath
 */
func (c *Config) DisablePreset(fileName string) error {
	err := os.Remove(c.EnabledPresetPath + fileName)
	if err != nil {
		return err
	}
	return nil
}