package require

import (
	"errors"
	"fmt"
	"testing"
)

func Test_NoError_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	var err error
	NoError(err)
}

func Test_NoError_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()

	err := errors.New("an error")
	NoError(err, "should not have an error")
}
