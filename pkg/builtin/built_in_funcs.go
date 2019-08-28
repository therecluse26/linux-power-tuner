package builtin

func Hello() string {
	return "Hello, World!"
}

func AddInts(num []int) int {
	total := 0
	for n := range num {
		total += n
	}
	return total
}