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

func getIrType(t graphite.Type) (string, error) {
	if t.IsPrimitive() {
		switch t.Name() {
		case "Int32":
			return "i32", nil
		default:
			return "", fmt.Errorf("could not serialize type %s", t.Name())
		}
	} else {
		return "", fmt.Errorf("dont know how to serialize non-primitive type %s", t.Name())
	}
}

func getInstructionName(m graphite.Method) (string, error) {
	irtype, err := getIrType(m.ReturnType())
	if err != nil {
		return "", err
	}
	switch m.Name() {
	case "+":
		return fmt.Sprintf("%s %s", "add", irtype), nil
	case "*":
		return fmt.Sprintf("%s %s", "mul", irtype), nil
	default:
		return "", fmt.Errorf("could not get instruction name for native method %s", m.Name())
	}
}
