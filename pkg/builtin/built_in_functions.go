package builtin

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func Notify(message interface{}) bool {
	fmt.Println(message)
	return true
}