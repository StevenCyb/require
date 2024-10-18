package require

import (
	"errors"
	"fmt"
	"runtime"
)

// ErrorIs asserts that an error is equal to a target error.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func ErrorIs(err, target error, msgAndArgs ...string) {
	if !errors.Is(err, target) {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expected error to be %v but got %v.\n",
			file, line, target, err)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
