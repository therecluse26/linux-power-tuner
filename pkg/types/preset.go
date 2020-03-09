package types

/*
 * Presets should be defined by pointing to events within "events" directory
 * within a .json preset file
 *
 * User-defined preset should be enabled/disabled by symlinking files
 * from "preset" into "enabled"
 */
type Preset struct {
	Name        	string 		`json:"name"`
	Description 	string 		`json:"description"`
	Hooks		   	[]Hook 		`json:"hooks"`
	Watchers 		[]Watcher 	`json:"watchers"`
}

/*
 * Go functions.
 */
type Function struct {
	Name			string		`json:"name"`
	Args []struct {
		Order 		int 		`json:"order"`
		Value 		interface{} `json:"value"`
	} `json:"args"`
}

/*
 * Event hooks. Wait for incoming input.
 */
type Hook struct {
	Name            string 		`json:"name"`
	Description     string 		`json:"description"`
	Slug			string		`json:"slug"`
	InputData		interface{}	`json:"input_data"`
	Reactions		[]Reaction 	`json:"reactions"`
}

/*
 * Monitors/watchers. Continuously poll endpoints.
 */
type Watcher struct {
	Name            string 		`json:"name"`
	Description     string 		`json:"description"`
	Interval 		int64  		`json:"interval"`
	Function		Function	`json:"function"`
	ExpectedVal 	interface{} `json:"expected_val"`
	Reactions   	[]Reaction 	`json:"reactions"`
}

type Reaction struct {
	Order			int			`json:"order"`
	Description		string		`json:"description"`
	Function 		Function	`json:"function"`
}