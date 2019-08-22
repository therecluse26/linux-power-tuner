package conditions

import (
	"fmt"
	"github.com/therecluse26/linux-power-tuner/pkg/utils"
	"os/exec"
	"regexp"
	"strings"
)

type CommandMeta struct {
	CommandString 	string
}


func (c *CommandMeta) Execute() ([]byte, error) {
	cmd := exec.Command("bash", "-c", c.CommandString)
	out, err := cmd.CombinedOutput()
	return out, err
}

func (c *CommandMeta) SearchCommandResult(search Search) bool {

	resultBuf, err := c.Execute()
	if err != nil {
		utils.HandleError(err, 0, true, true)
		return false
	}

	if search.Type == Simple {
		if fmt.Sprintf("%v", search.Query) == strings.TrimSpace(string(resultBuf)) {
			return true
		}
	} else if search.Type == Regex {

		match, err := regexp.Match(search.Query.(string), resultBuf)
		if err != nil {
			utils.HandleError(err, 0, true, true)
			return false
		}
		if match {
			return true
		}
	}

	return false

}
