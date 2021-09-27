package tools

import "log"

// FailOnError return generic error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
	}
}