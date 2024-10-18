package require

import (
	"errors"
	"fmt"
	"runtime"
)

// ErrorAs asserts that an error can be assigned to a target type.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func ErrorAs(err error, target any, msgAndArgs ...string) {
	if !errors.As(err, target) {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expected error to be assignable to %T but it is not.\n",
			file, line, target)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
