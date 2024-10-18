package require

import (
	"fmt"
	"testing"
)

func Test_NotNil_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	a := new(int)
	NotNil(a)
}

func Test_NotNil_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()

	var a *int
	NotNil(a, "should not be nil")
}
