package utils_test

import (
	"testing"

	utils "github.com/CoachApplication/utils"
)

type TestStruct struct {
	id      string
	message string
}

func (t TestStruct) Id() string {
	return t.id
}

func (t TestStruct) Message() string {
	return t.message
}

func TestStandardOrderedMap_Get(t *testing.T) {
	m := utils.StandardOrderedMap{}
	val := TestStruct{id: "one", message: "this is one"}

	m.Add(val)
	if getVal, err := m.Get("one"); err != nil {
		t.Error("StandardOrderedMap returned an error when retrieving a valid id: ", err)
	} else if convVal, good := getVal.(TestStruct); !good {
		t.Error("StandardOrderedMap returned an item of the wrong type")
	} else if convVal.Message() != "this is one" {
		t.Error("StandardOrderedMap retrieved item returned the wrong message: ", convVal.Message())
	}

	if _, err := m.Get("two"); err == nil {
		t.Error("StandardOrderedMap did not return an error when asked for an invalid key")
	}
}

func TestStandardOrderedMap_Add(t *testing.T) {
	m := utils.StandardOrderedMap{}
	val := TestStruct{id: "one", message: "this is one"}

	m.Add(val)
	if getVal, err := m.Get("one"); err != nil {
		t.Error("StandardOrderedMap returned an error when retrieving a valid id: ", err)
	} else if convVal, good := getVal.(TestStruct); !good {
		t.Error("StandardOrderedMap returned an item of the wrong type")
	} else if convVal.Message() != "this is one" {
		t.Error("StandardOrderedMap retrieved item returned the wrong message: ", convVal.Message())
	}

}

func TestStandardOrderedMap_Set(t *testing.T) {
	m := utils.StandardOrderedMap{}
	val := TestStruct{id: "one", message: "this is one"}

	m.Set("one", val)
	if getVal, err := m.Get("one"); err != nil {
		t.Error("StandardOrderedMap returned an error when retrieving a valid id: ", err)
	} else if convVal, good := getVal.(TestStruct); !good {
		t.Error("StandardOrderedMap returned an item of the wrong type")
	} else if convVal.Message() != "this is one" {
		t.Error("StandardOrderedMap retrieved item returned the wrong message: ", convVal.Message())
	}
}

func TestStandardOrderedMap_Order(t *testing.T) {

	m := utils.StandardOrderedMap{}

	m.Add(TestStruct{id: "one", message: "this is one"})
	m.Add(TestStruct{id: "two", message: "this is two"})
	m.Add(TestStruct{id: "three", message: "this is three"})

	l := m.Order()
	if len(l) != 3 {
		t.Error("StandardOrderedMap Order returned the wrong number of args")
	} else if l[0] != "one" || l[1] != "two" || l[2] != "three" {
		t.Error("StandardOrderedMap Order did not return the correct order")
	}
}
