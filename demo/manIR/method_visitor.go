package manIR

import (
	"github.com/mniak/graphite"
	"github.com/pkg/errors"
	"strconv"
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
		v.WriteString(" %")
		v.WriteString(param.Name())
	}
	v.WriteString(") {\n")
	v.Indent()

	err := m.Body().AcceptValueVisitor(v)

	v.Dedent()
	v.WriteString("}\n")

	return err
}

func (v *methodVisitor) VisitNativeOperator(o graphite.NativeOperator) error {
	v.WriteString("<native operator/>\n")
	return nil
}

func (v *methodVisitor) VisitInvocation(mi graphite.Invocation) error {
	args := mi.Arguments()
	//v.WriteString("(")
	method := mi.Method()
	varname := v.newvar()
	v.WriteString(varname)
	v.WriteString(" = ")
	if method.IsNative() {
		instr, err := getInstructionName(method)
		if err != nil {
			return errors.Wrap(err, "failed to serialize native instruction name")
		}
		v.WriteString(instr)
		v.WriteString(" ")
		v.serializeArgumentList(args, false)
	} else {
		v.WriteString("call ")
		v.WriteString(mi.Method().Name())
		v.WriteString(" ")
		v.serializeArgumentList(args, true)
	}
	v.WriteString("\n")
	return nil
}
func (v *methodVisitor) newvar() string {
	return "%v_0"
}
func (v *methodVisitor) serializeArgumentList(args []graphite.Argument, withParenthesis bool) error {
	first := true
	if withParenthesis {
		v.WriteString("(")
	}
	for _, arg := range args {
		if first {
			first = false
		} else {
			v.WriteString(" ")
		}
		err := arg.Value().AcceptValueVisitor(v)
		if err != nil {
			return errors.Wrap(err, "failed to serialize argument")
		}
	}
	if withParenthesis {
		v.WriteString(")")
	}
	return nil
}

func (v *methodVisitor) VisitInt32Literal(i int32) error {
	v.WriteString(strconv.Itoa(int(i)))
	return nil
}

func (v *methodVisitor) VisitParameterValue(pv graphite.ParameterValue) error {
	v.WriteString(pv.Parameter().Name())
	return nil
}
