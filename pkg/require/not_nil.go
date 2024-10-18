package require

import (
	"fmt"
	"reflect"
	"runtime"
)

// NotNil asserts that the specified object is not nil.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func NotNil[T any](actual *T, msgAndArgs ...string) {
	if reflect.ValueOf(actual).IsNil() {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expected value not to be nil.\n",
			file, line)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
