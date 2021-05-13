package ir

import "fmt"

type scope struct {
	variables map[string]int
}

func newScope() scope {
	return scope{
		variables: make(map[string]int, 0),
	}
}

func (s *scope) findFreeName(nameHint string) (string, int) {
	n, has := s.variables[nameHint]
	if !has {
		return nameHint, 1
	}
	return s.findFreeName(fmt.Sprintf("%s%d", nameHint, n))
}

func (s *scope) AddVariable(nameHint string) string {
	name, n := s.findFreeName(nameHint)
	s.variables[name] = n
	return name
}
