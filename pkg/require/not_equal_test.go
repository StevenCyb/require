package require

import (
	"fmt"
	"testing"

	"github.com/StevenCyb/require/pkg/require/internal"
)

func Test_NotEqual_Number_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()
	NotEqual(1, 2)
}

func Test_NotEqual_Number_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	NotEqual(1, 1, "Numbers are equal")
}

func Test_NotEqual_String_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	NotEqual("hello", "world")
}

func Test_NotEqual_String_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	NotEqual("hello", "hello")
}

func Test_NotEqual_Array_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	NotEqual([3]int{1, 2, 3}, [3]int{4, 5, 6})
}

func Test_NotEqual_Array_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	NotEqual([3]int{1, 2, 3}, [3]int{1, 2, 3})
}

func Test_NotEqual_Map_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	NotEqual(map[string]int{"one": 1, "two": 2}, map[string]int{"three": 3, "four": 4})
}

func Test_NotEqual_Map_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	NotEqual(map[string]int{"one": 1, "two": 2}, map[string]int{"one": 1, "two": 2})
}

func Test_NotEqual_Object_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()

	NotEqual(internal.Person{Name: "John", Age: 30}, internal.Person{Name: "Jane", Age: 25})
}

func Test_NotEqual_Object_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	NotEqual(internal.Person{Name: "John", Age: 30}, internal.Person{Name: "John", Age: 30})
}
