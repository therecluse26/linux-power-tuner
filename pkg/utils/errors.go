package utils

import "log"

const (
	Panic = iota
	Error
	Warning
	Info
)

// Main error handler
func HandleError (e error, severity int, log bool, display bool) {
	if log == true {
		LogError(e, severity)
	}
	if display == true {
		DisplayError(e, severity)
	}
}

// Displays error in GUI
func DisplayError(e error, severity int) {

	// Display error code here

}

// Logs error
func LogError(e error, severity int) {
	if severity == Panic {
		log.Fatal(e)
	} else {
		log.Println(e)
	}
}
