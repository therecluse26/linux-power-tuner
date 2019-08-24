package function

type Arguments []struct {
		Key   string `json:"key"`
		Value interface{} `json:"value"`
	}

type Function struct {
	Name string `json:"name"`
	Args Arguments `json:"args"`
}