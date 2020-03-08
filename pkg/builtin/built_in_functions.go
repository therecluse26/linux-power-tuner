package builtin

import (
	"github.com/gen2brain/beeep"
)

type NotificationType string

func Hello() string {
	return "Hello, World!"
}

func Notify(title interface{}, message interface{}) bool {

	beeep.Alert(title.(string),message.(string), "")
	return true
}