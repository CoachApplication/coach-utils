package utils

import (
	"fmt"
)

type OrderedMap interface {
	Set(key string, val interface{}) error
	Get(key string) (interface{}, error)
	Order() []string
}

/**
 * OrderedHasIdMap an ordered map where each value can identify itself
 */
type OrderedHasIdMap interface {
	Add(val HasId) error
	Get(key string) (interface{}, error)
	Order() []string
}

// HasId can identify itself
type HasId interface {
	Id() string
}

// StandardOrderedMap
type StandardOrderedMap struct {
	vMap   map[string]interface{}
	vOrder []string
}

func NewOrderedMap() OrderedMap {
	return OrderedMap(&StandardOrderedMap{})
}

func NewOrderedHasIdMap() OrderedHasIdMap {
	return OrderedHasIdMap(&StandardOrderedMap{})
}

func (som *StandardOrderedMap) safe() {
	if som.vMap == nil {
		som.vMap = map[string]interface{}{}
		som.vOrder = []string{}
	}
}

func (som *StandardOrderedMap) Set(id string, val interface{}) error {
	som.safe()
	if _, found := som.vMap[id]; !found {
		som.vOrder = append(som.vOrder, id)
	}
	som.vMap[id] = val
	return nil
}

func (som *StandardOrderedMap) Add(val HasId) error {
	id := val.Id()
	return som.Set(id, val)
}

func (som *StandardOrderedMap) Get(id string) (interface{}, error) {
	som.safe()
	if val, found := som.vMap[id]; found {
		return val, nil
	} else {
		return val, error(IdNotFoundError{id: id})
	}
}

func (som *StandardOrderedMap) Order() []string {
	som.safe()
	return som.vOrder
}

// IdNotFoundError An error for when the Map was asked for an Id that has not been Add()/Set()
type IdNotFoundError struct {
	id string
}

// Error return an error string (interface: error)
func (infe IdNotFoundError) Error() string {
	return fmt.Sprintf("Id %s was not found in the ordered map", infe.id)
}
