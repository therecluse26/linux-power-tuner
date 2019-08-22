package conditions

type Search struct {
	Type SearchType
	Query interface{}
}

type SearchType int

const (
	Simple SearchType = iota
	Regex
	Csv
)
