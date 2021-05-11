package manIR

import (
	fmt "fmt"
	"github.com/mniak/graphite"
	"github.com/pkg/errors"
)

type methodVisitor struct {
	parent *programVisitor
}

func (v *methodVisitor) WriteString(str string) {
	v.parent.sb.WriteString(str)
}

func (v *methodVisitor) Indent() {
	v.parent.sb.Indent()
}

func (v *methodVisitor) Dedent() {
	v.parent.sb.Dedent()
}

func (v *methodVisitor) VisitInternalMethod(m graphite.InternalMethod) error {
	irReturnType, err := getIrType(m.ReturnType())
	if err != nil {
		return errors.Wrap(err, "error serializing method return type")
	}
	v.WriteString(fmt.Sprintf("define %s %s(", irReturnType, formatFunctionName(m.Name())))
	params := m.Parameters()
	first := true
	for _, param := range params {
		if first {
			first = false
		} else {
			v.WriteString(", ")
		}
		irType, err := getIrType(param.ReturnType())
		if err != nil {
			return errors.Wrap(err, "error serializing parameter return type")
		}
		v.WriteString(fmt.Sprintf("%s %s", irType, formatParameterName(param.Name())))
	}
	v.WriteString(") {\n")
	v.Indent()

	valueVisitor := valueVisitor{
		parent: v.parent,
	}
	err = m.Body().AcceptValueVisitor(&valueVisitor)
	if err != nil {
		return errors.Wrap(err, "error serializing value")
	}
	v.WriteString(fmt.Sprintf("ret %s %s\n", irReturnType, valueVisitor.lastExpression))

	v.Dedent()
	v.WriteString("}\n")
	return err
}

func (v *methodVisitor) VisitNativeOperation(graphite.NativeOperation) error {
	return nil
}
