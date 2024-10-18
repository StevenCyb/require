package require

import (
	"errors"
	"fmt"
	"testing"

	"github.com/StevenCyb/require/pkg/require/internal"
)

func Test_ErrorAs_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	var target *internal.CustomError
	err := &internal.CustomError{}
	ErrorAs(err, &target)
}

func Test_ErrorAs_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()

	var target *internal.CustomError
	err := errors.New("an error")
	ErrorAs(err, &target, "error should be assignable to target type")
}
