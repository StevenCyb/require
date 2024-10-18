package require

import (
	"fmt"
	"reflect"
	"runtime"
)

// Equal asserts that two objects are equal.
// Optionally, a message can be passed as the final argument (only first string is interpreted).
func Equal[T any](expect T, actual T, msgAndArgs ...string) {
	if !reflect.DeepEqual(expect, actual) {
		_, file, line, _ := runtime.Caller(1)
		errorMessage := fmt.Sprintf("--- FAIL:\n"+
			"    %s:%d:\n"+
			"          Error: Not equal:\n"+
			"                 expected: %v\n"+
			"                 actual  : %v\n",
			file, line, expect, actual)

		if len(msgAndArgs) > 0 {
			errorMessage = fmt.Sprintf("%s                 Messages: %s\n", errorMessage, msgAndArgs[0])
		}

		panic(errorMessage)
	}
}
