package require

import (
	"fmt"
	"reflect"
	"runtime"
)

// Nil asserts that the specified parameter is nil.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func Nil[T any](actual *T, msgAndArgs ...string) {
	if !reflect.ValueOf(actual).IsNil() {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Expect nil, but got: %v\n",
			file, line, actual)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
