package utils

import "log"

// CheckAndPanic ...
func CheckAndPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// CheckAndLogFatal ...
func CheckAndLogFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// CheckAndLog ...
func CheckAndLog(e error) {
	if e != nil {
		log.Print(e)
	}
}
