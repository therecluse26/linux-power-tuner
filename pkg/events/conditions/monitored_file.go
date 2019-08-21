package conditions

import (
	"fmt"
	"github.com/therecluse26/linux-power-tuner/pkg/utils"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"os"
	"strings"
)

type FileMeta struct {
	Path        string
	AccessLevel FileAccess
	Opened 		bool
}

type Search struct {
	Type SearchType
	Query interface{}
}

type FileAccess int

type SearchType int

const (
	Writeable FileAccess = iota
	Readable
	Unreadable
	NotFound
)

const (
	Simple SearchType = iota
	Regex
	Csv
)

/*
 * Checks access level of a file path
 */
func (f *FileMeta) CheckAccess() {

	// If file not found
	_, fileErr := os.Stat(f.Path)
	if os.IsNotExist(fileErr) {
		utils.HandleError(fileErr, 1, true, true)
		f.AccessLevel = NotFound
		return
	}

	// If file is writeable
	fileErr = unix.Access(f.Path, unix.W_OK)
	if fileErr == nil  {
		f.AccessLevel = Writeable
		return
	}

	// If file is readable
	fileErr = unix.Access(f.Path, unix.R_OK)
	if fileErr == nil {
		f.AccessLevel = Readable
		return
	} else {
		utils.HandleError(fileErr, 1, true, true)
		f.AccessLevel = Unreadable
		return
	}

}

func (f *FileMeta) SearchFileValue(search Search) bool {

	f.CheckAccess()

	if f.AccessLevel == Readable || f.AccessLevel == Writeable {

		if search.Type == Simple {

			buf, err := ioutil.ReadFile(f.Path)
			if err != nil {
				utils.HandleError(err, 0, true, true)
				return false
			}

			if fmt.Sprintf("%v", search.Query) == strings.TrimSpace(string(buf)) {
				return true
			}

		} else if search.Type == Regex {



		}

	}

	return false

}