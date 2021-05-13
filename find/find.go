package find

import (
	"fmt"
	"reflect"

	"github.com/mniak/graphite"
)

func Find(node interface{}, match func(interface{}) bool) ([]interface{}, error) {
	visitor := findVisitor{
		match:   match,
		results: make([]interface{}, 0),
	}
	var err error
	switch x := node.(type) {
	case graphite.Program:
		err = visitor.VisitProgram(x)
	case graphite.Invocation:
		err = visitor.VisitInvocation(x)
	case graphite.ParameterValue:
		err = visitor.VisitParameterValue(x)
	case graphite.Argument:
		err = visitor.findInArgument(x)
	default:
		return visitor.results, fmt.Errorf("could not traverse %v", x)
	}
	return visitor.results, err
}

func Type(root interface{}, typee reflect.Type) (result []interface{}, err error) {
	result, err = Find(root, func(node interface{}) bool {
		r := reflect.TypeOf(node).Implements(typee)
		return r
	})
	return
}
