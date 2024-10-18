package require

import (
	"fmt"
	"testing"
)

func Test_Nil_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	var a *int
	Nil(a)
}

func Test_Nil_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()

	a := new(int)
	Nil(a, "should be nil")
}
