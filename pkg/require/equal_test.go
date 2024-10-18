package require

import (
	"fmt"
	"testing"

	"github.com/StevenCyb/require/pkg/require/internal"
)

func Test_Equal_Number_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()
	Equal(1, 1)
}

func Test_Equal_Number_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	Equal(1, 2, "Some message")
}

func Test_Equal_String_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()
	Equal("hello", "hello")
}

func Test_Equal_String_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	Equal("hello", "world", "Strings are not equal")
}

func Test_Equal_Array_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()
	Equal([3]int{1, 2, 3}, [3]int{1, 2, 3})
}

func Test_Equal_Array_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	Equal([3]int{1, 2, 3}, [3]int{4, 5, 6}, "Arrays are not equal")
}

func Test_Equal_Map_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()
	Equal(map[string]int{"one": 1, "two": 2}, map[string]int{"one": 1, "two": 2})
}

func Test_Equal_Map_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	Equal(map[string]int{"one": 1, "two": 2}, map[string]int{"three": 3, "four": 4}, "Maps are not equal")
}

func Test_Equal_Object_Success(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Test failed: %v", r)
		}
	}()
	Equal(internal.Person{Name: "John", Age: 30}, internal.Person{Name: "John", Age: 30})
}

func Test_Equal_Object_Failure(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test did not fail as expected")
		} else {
			fmt.Println(r)
		}
	}()
	Equal(internal.Person{Name: "John", Age: 30}, internal.Person{Name: "Jane", Age: 25}, "Objects are not equal")
}
