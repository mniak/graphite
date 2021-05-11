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
	v.WriteString("define ")
	v.WriteString("@")
	v.WriteString(m.Name())
	v.WriteString("(")
	params := m.Parameters()
	first := true
	for _, param := range params {
		if first {
			first = false
		} else {
			v.WriteString(", ")
		}
		err := serializeIrType(v, param.ReturnType())
		if err != nil {
			return errors.Wrap(err, "error serializing parameter return type")
		}
		v.WriteString(fmt.Sprintf(" %s", formatParameterName(param.Name())))
	}
	v.WriteString(") {\n")
	v.Indent()

	valueVisitor := valueVisitor{
		parent: v.parent,
	}
	err := m.Body().AcceptValueVisitor(&valueVisitor)
	v.WriteString(fmt.Sprintf("ret %s\n", valueVisitor.lastExpression))

	v.Dedent()
	v.WriteString("}\n")
	return err
}

func (v *methodVisitor) VisitNativeOperation(m graphite.NativeOperation) error {
	return nil
}
