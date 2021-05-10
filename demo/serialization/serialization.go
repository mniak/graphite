package serialization

import "github.com/mniak/graphite"

func Serialize(program graphite.Program) (string, error) {
	var visitor visitor
	err := program.Accept(&visitor)
	return visitor.String(), err
}
