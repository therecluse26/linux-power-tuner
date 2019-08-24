package function


type Function struct {
	Name string `json:"name"`
	Args []struct {
		Key   string `json:"key"`
		Value interface{} `json:"value"`
	} `json:"args"`
}