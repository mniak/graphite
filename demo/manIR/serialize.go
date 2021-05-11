package manIR

import (
	"fmt"
	"github.com/mniak/graphite"
)

func SerializeProgram(program graphite.Program) (string, error) {
	var visitor programVisitor
	err := visitor.serializeProgram(program)
	return visitor.String(), err
}

type writer interface {
	WriteString(string)
}

func serializeIrType(w writer, t graphite.Type) error {
	if t.IsPrimitive() {
		switch t.Name() {
		case "Int32":
			w.WriteString("i32")
		default:
			return fmt.Errorf("could not serialize type %s", t.Name())
		}
	} else {
		return fmt.Errorf("dont know how to serialize non-primitive type %s", t.Name())
	}
	return nil
}

func getInstructionName(m graphite.Method) (string, error) {
	switch m.Name() {
	case "+":
		return "add", nil
	case "*":
		return "mul", nil
	default:
		return "", fmt.Errorf("could not get instruction name for native method %s", m.Name())
	}
}
