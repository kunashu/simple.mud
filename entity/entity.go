package entity

import (
	"bytes"
	"strings"
)

type EntityId uint

type Entity interface {
	Name() string
	Id() EntityId
	FullMatch(name string) bool
	Match(name string) bool
	Data() interface{}
}

type baseEntity struct {
	Name string
	Id   EntityId
}

func (e baseEntity) CompName() string {
	return strings.ToLower(e.name)
}

func (e baseEntity) FullMatch(name string) bool {
	return strings.EqualFold(e.CompName(), name)
}

func (e baseEntity) Match(name string) bool {
	if len(name) == 0 {
		return true
	}

	e_name := []byte(e.CompName())
	pos := bytes.IndexAny(e_name, name)
	for pos != -1 {
		if pos == 0 || e_name[pos-1] == ' ' {
			return true
		}

		e_name = e_name[pos+1:]
		pos = bytes.IndexAny(e_name, name)
	}

	return false
}
