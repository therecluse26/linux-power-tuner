package builtin

import (
	"bytes"
	"github.com/gen2brain/beeep"
	"github.com/spf13/viper"
	"os/exec"
)

type NotificationType string

func Hello() string {
	return "Hello, World!"
}

/*
 * Dispatch desktop notification
 */
func Notify(title interface{}, message interface{}) bool {
	err := beeep.Alert(title.(string),message.(string), viper.Get("ProjectLogo").(string))
	if err != nil {
		return false
	}
	return true
}

/*
 * Execute Shell Command
 */
func ShellCommand(name interface{}, arg ...interface{}) (string, error) {
	var strArg []string
	for _, a := range arg {
		strArg = append(strArg, a.(string))
	}
	cmd := exec.Command(name.(string), strArg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}