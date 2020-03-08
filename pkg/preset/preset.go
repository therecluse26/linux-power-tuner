package preset

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/therecluse26/uranium/pkg/function"
	"github.com/therecluse26/uranium/pkg/types"
	"github.com/therecluse26/uranium/pkg/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)


func RunPresets(presets []types.Preset, loadedFunctions types.Funcs){
	for _, pre := range presets {
		for _, ev := range pre.Events {
			for _, fn := range ev.Conditions {
				res, err := function.CallFunction(loadedFunctions, fn.Function)
				if err != nil {
					utils.HandleError(err, 1, true, true)
				}
				fmt.Println(res)
			}
		}
	}
}
/*
 * Loads active preset files from EnabledPresetPath directory
 */
func LoadActivePresets() []types.Preset {
	var presets []types.Preset
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
func LoadPreset(filePath string) types.Preset {
	presetFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		utils.HandleError(err, 0, true, true)
	}
	presetStruct := types.Preset{}
	err = json.Unmarshal(presetFile, &presetStruct)
	return presetStruct
}

/*
 * Gets available preset file names from AvailablePresetPath
 */
func GetAvailablePresets() []string {
	var presetFiles []string
	_ = filepath.Walk( viper.Get("AvailablePresetPath").(string), func(path string, info os.FileInfo, err error) error {

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
	err := os.Link( viper.Get("AvailablePresetPath").(string) + fileName,
					viper.Get("EnabledPresetPath").(string) + fileName )
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
