package ir_manual

import (
	fmt "fmt"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/render/writer"
	"github.com/pkg/errors"
)

type methodVisitor struct {
	writer writer.Writer
}

func (v *methodVisitor) VisitInternalMethod(m graphite.InternalMethod) error {
	irReturnType, err := getIrType(m.ReturnType())
	if err != nil {
		return errors.Wrap(err, "error serializing method return type")
	}
	v.writer.WriteString(fmt.Sprintf("define %s %s(", irReturnType, formatFunctionName(m.Name())))
	params := m.Parameters()
	first := true
	for _, param := range params {
		if first {
			first = false
		} else {
			v.writer.WriteString(", ")
		}
		irType, err := getIrType(param.ReturnType())
		if err != nil {
			return errors.Wrap(err, "error serializing parameter return type")
		}
		v.writer.WriteString(fmt.Sprintf("%s %s", irType, formatParameterName(param.Name())))
	}
	v.writer.WriteString(") {\n")
	v.writer.Indent()

	valueVisitor := valueVisitor{
		writer: v.writer,
	}
	err = m.Body().AcceptValueVisitor(&valueVisitor)
	if err != nil {
		return errors.Wrap(err, "error serializing value")
	}
	v.writer.WriteString(fmt.Sprintf("ret %s %s\n", irReturnType, valueVisitor.lastExpression))

	v.writer.Dedent()
	v.writer.WriteString("}\n")
	return err
}

func (v *methodVisitor) VisitNativeOperation(graphite.NativeOperation) error {
	return nil
}
