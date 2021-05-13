package manualIR

import (
	"fmt"

	"github.com/mniak/graphite"
	"github.com/mniak/graphite/find"
	"github.com/mniak/graphite/render/writer"
	"github.com/pkg/errors"
)

func SerializeProgram(program graphite.Program) (string, error) {
	w := writer.New()
	methods, err := find.Methods(program)
	if err != nil {
		return "", errors.Wrap(err, "error finding methods")
	}

	for _, method := range methods {
		methodVisitor := methodVisitor{
			writer: w,
		}
		err := method.AcceptMethodVisitor(&methodVisitor)
		if err != nil {
			return "", errors.Wrap(err, "error serializing method")
		}

	}

	w.WriteString("\ndefine i32 @main() {\n")
	w.Indent()
	valueVisitor := valueVisitor{
		writer: w,
	}
	err = program.Entrypoint().AcceptValueVisitor(&valueVisitor)
	w.WriteString(fmt.Sprintf("ret i32 %s\n", valueVisitor.lastExpression))
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize statement")
	}

	w.Dedent()
	w.WriteString("}\n")
	return w.String(), nil
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
	irtype, err := getIrType(m.Type())
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
