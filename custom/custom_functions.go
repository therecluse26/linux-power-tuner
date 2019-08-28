package custom

import "github.com/gobuffalo/flect"

func UpperCase(i string) string {
	return flect.Capitalize(i)
}

func Multiply(a, b, c int) int {
	return a * b * c
}
