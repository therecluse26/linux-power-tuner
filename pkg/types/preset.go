package types

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
		Name            string `json:"name"`
		Description     string `json:"description"`
		PollingInterval int64  `json:"pollingInterval"`
		Conditions      []struct {
			Id 			int    		`json:"id"`
			Description string 		`json:"description"`
			Function 	Function		`json:"function"`
			ExpectedVal interface{} `json:"expected_val"`
		} `json:"conditions"`
		ConditionExp string `json:"condition_exp"`
	} `json:"events"`
	Reactions   []struct {
		Name     string		`json:"name"`
		Function Function	`json:"function"`
	} `json:"reactions"`
}

type Function struct {
	Name string `json:"name"`
	Args []struct {
		Order int `json:"order"`
		Value interface{} `json:"value"`
	} `json:"args"`
}

