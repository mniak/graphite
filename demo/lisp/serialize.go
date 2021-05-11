package lisp

import "github.com/mniak/graphite"

func Serialize(program graphite.Program) (string, error) {
	var visitor visitor
	err := visitor.serializeProgram(program)
	return visitor.String(), err
}
