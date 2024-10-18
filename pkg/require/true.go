package require

import (
	"fmt"
	"runtime"
)

// True asserts that the specified value is true.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func True(value bool, msgAndArgs ...string) {
	if !value {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expected value to be true but got false.\n",
			file, line)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
