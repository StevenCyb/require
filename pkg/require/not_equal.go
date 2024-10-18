package require

import (
	"fmt"
	"reflect"
	"runtime"
)

// NotEqual asserts that the specified values are NOT equal.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func NotEqual[T any](expect T, actual T, msgAndArgs ...string) {
	if reflect.DeepEqual(expect, actual) {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Should not be: %v\n",
			file, line, actual)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
