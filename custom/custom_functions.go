package custom

import (
	"strings"
)

func UpperCase(i interface{}) string {
	return strings.ToUpper(i.(string))
}

func AddNumbers(nums []interface{}) float64 {
	var total float64
	for _, n := range nums {
		total += n.(float64)
	}
	return total
}

func Multiply(nums []interface{}) float64 {
	var t = nums[0].(float64)
	for i := 0; i < len(nums); i++ {
		t *= nums[i].(float64)
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