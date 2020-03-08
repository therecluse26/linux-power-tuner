package builtin

import (
	"github.com/gen2brain/beeep"
)

type NotificationType string

const (
	NotificationNormal NotificationType = "normal"
	NotificationSuccess = "success"
	NotificationError = "error"
)

func Hello() string {
	return "Hello, World!"
}

func Notify(message interface{}, notifyType ...interface{}) bool {
	switch NotificationType(notifyType[0].(string)) {
	case NotificationNormal:
		beeep.Alert("" ,message.(string), "")
	case NotificationSuccess:
		beeep.Alert("Uranium Success" ,message.(string), "")
	case NotificationError:
		beeep.Alert("Uranium Error" ,message.(string), "")
		return false
	default:
		beeep.Alert("" ,message.(string), "")
	}
	return true
}