package preset

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/therecluse26/uranium/pkg/function"
	"github.com/therecluse26/uranium/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
 * Presets should be defined by pointing to events within "events" directory
 * within a .json preset file
 *
 * User-defined preset should be enabled/disabled by symlinking files
 * from "preset" into "enabled"
 */

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
			Function 	function.Function	`json:"function"`
			ExpectedVal interface{} `json:"expected_val"`
		} `json:"conditions"`
		ConditionExp string `json:"condition_exp"`
	} `json:"events"`
	Reactions   []struct {
		Name     string		`json:"name"`
		Function function.Function	`json:"function"`
	} `json:"reactions"`
}

/*
 * Loads active preset files from EnabledPresetPath directory
 */
func LoadActivePresets() []Preset {
	var presets []Preset
	_ = filepath.Walk(viper.Get("EnabledPresetPath").(string), func(path string, info os.FileInfo, err error) error {
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
func GetAvailablePresets() []string {
	var presetFiles []string
	_ = filepath.Walk(viper.Get("AvailablePresetPath").(string), func(path string, info os.FileInfo, err error) error {
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
func EnablePreset(fileName string) error {
	err := os.Link(viper.Get("AvailablePresetPath").(string) + fileName, viper.Get("EnabledPresetPath").(string) + fileName)
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
func DisablePreset(fileName string) error {
	err := os.Remove(viper.Get("EnabledPresetPath").(string) + fileName)
	if err != nil {
		return err
	}
	return nil
}
