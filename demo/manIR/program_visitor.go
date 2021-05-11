package manIR

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/find"
	"github.com/pkg/errors"
)

type programVisitor struct {
	sb indentedStringBuilder
}

func (v *programVisitor) String() string {
	return v.sb.String()
}

func (v *programVisitor) serializeProgram(program graphite.Program) error {
	methods, err := find.Methods(program)
	if err != nil {
		return errors.Wrap(err, "error finding methods")
	}

	for _, method := range methods {
		methodVisitor := methodVisitor{
			parent: v,
		}
		err := method.AcceptMethodVisitor(&methodVisitor)
		if err != nil {
			return errors.Wrap(err, "error serializing method")
		}

	}

	v.sb.WriteString("\ndefine i32 @main() {\n")
	v.sb.Indent()
	valueVisitor := valueVisitor{
		parent: v,
	}
	err = program.Entrypoint().AcceptValueVisitor(&valueVisitor)
	v.sb.WriteString(fmt.Sprintf("ret i32 %s\n", valueVisitor.lastExpression))
	if err != nil {
		return errors.Wrap(err, "failed to serialize statement")
	}

	v.sb.Dedent()
	v.sb.WriteString("}")
	return nil
}
