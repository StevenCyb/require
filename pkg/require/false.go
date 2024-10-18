package require

import (
	"fmt"
	"runtime"
)

// False asserts that the specified value is false.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func False(value bool, msgAndArgs ...string) {
	if value {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expected value to be false but got true.\n",
			file, line)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
