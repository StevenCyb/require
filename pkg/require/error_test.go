package require

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Error_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	err := errors.New("an error")
	Error(err)
}

func Test_Error_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()

	var err error
	Error(err, "should have an error")
}
