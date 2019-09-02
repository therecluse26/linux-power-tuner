package builtin

func Hello() string {
	return "Hello, World!"
}

func AddInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}