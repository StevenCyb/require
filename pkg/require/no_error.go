package require

import (
	"fmt"
	"runtime"
)

// NoError asserts that an error is nil.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func NoError(err error, msgAndArgs ...string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expected no error but got: %v\n",
			file, line, err)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
