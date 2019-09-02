package custom

import (
	"strings"
)

func UpperCase(i string) string {
	return strings.ToUpper(i)
}

func Multiply(nums []int) int {
	var t = nums[0]
	for i := 0; i < len(nums); i++ {
		t *= nums[i]
	}
	return t
}

func Map(slice []interface{}, function func(interface{})interface{}) interface{} {
	var newSlice []interface{}
	for _, n := range slice {
		newSlice = append(newSlice, function(n))
	}
	return newSlice
}