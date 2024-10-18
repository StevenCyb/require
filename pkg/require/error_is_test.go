package require

import (
	"errors"
	"fmt"
	"testing"
)

var (
	errAnError      = errors.New("an error")
	errAnotherError = errors.New("another error")
)

func Test_ErrorIs_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	target := errAnError
	ErrorIs(errAnError, target)
}

func Test_ErrorIs_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()

	err := errors.New("an error")
	target := errors.New("another error")
	ErrorIs(err, target, "errors should match")
}
