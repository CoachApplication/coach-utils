package utils_test

import (
	"testing"

	utils "github.com/CoachApplication/coach-utils"
)

// We use the unique string slice in the backend provider, so we should test that it works
func Test_UniqueStringSlice(t *testing.T) {
	s := utils.UniqueStringSlice{}

	// test adding strings
	s.Add("A")
	sl1 := s.Slice()
	if len(sl1) != 1 {
		t.Error("uniqueStringSlice did not properly add a value")
	} else if sl1[0] != "A" {
		t.Error("uniqueStringSlice did not properly add a value")
	}

	// test adding further
	s.Add("B")
	s.Add("C")
	sl2 := s.Slice()
	if len(sl2) != 3 {
		t.Error("uniqueStringSlice did not properly add a value")
	} else if sl2[0] != "A" || sl2[1] != "B" || sl2[2] != "C" {
		t.Error("uniqueStringSlice did not properly add a value")
	}

	// test unique values
	s.Add("A")
	sl3 := s.Slice()
	if len(sl3) != 3 {
		t.Error("uniqueStringSlice improperly added a duplicate value")
	}
}
