package require

import (
	"fmt"
	"runtime"
)

// Error asserts that an error is not nil.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func Error(err error, msgAndArgs ...string) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expected an error but got nil.\n",
			file, line)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
