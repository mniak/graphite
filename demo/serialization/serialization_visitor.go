package serialization

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/find"
	"github.com/pkg/errors"
)

type visitor struct {
	sb indentedStringBuilder
}

func (v *visitor) String() string {
	return v.sb.String()
}

func (v *visitor) VisitInternalMethod(m graphite.InternalMethod) error {
	v.sb.WriteString(fmt.Sprintf("DECL %s (", m.Name()))
	for _, param := range m.Parameters() {
		v.serializeParameter(param)
	}
	v.sb.WriteString(fmt.Sprintf(") -> %s\n", m.ReturnType().Name()))
	v.sb.Indent()
	err := m.Body().AcceptValueVisitor(v)
	v.sb.Dedent()
	return err
}

func (v *visitor) VisitInvocation(mi graphite.Invocation) error {
	args := mi.Arguments()
	v.sb.WriteString(fmt.Sprintf("INVOKE %s (", mi.Method().Name()))
	for _, arg := range args {
		err := v.serializeArgument(arg)
		if err != nil {
			return errors.Wrap(err, "failed to serialize argument")
		}
	}
	v.sb.WriteString(fmt.Sprintf(") -> %s", mi.Method().ReturnType().Name()))
	return nil
}

func (v *visitor) VisitInt32Literal(i int32) error {
	v.sb.WriteString(fmt.Sprintf("[int32] %d", i))
	return nil
}

func (v *visitor) VisitParameterValue(pv graphite.ParameterValue) error {
	v.sb.WriteString(fmt.Sprintf("[param] %s", pv.Parameter().Name()))
	return nil
}

func (v *visitor) serializeArgument(a graphite.Argument) error {
	v.sb.WriteString(fmt.Sprintf("{%s = ", a.Parameter().Name()))
	defer v.sb.WriteString("}")
	return a.Value().AcceptValueVisitor(v)
}
func (v *visitor) serializeProgram(program graphite.Program) error {
	v.sb.WriteString(":DECLARATIONS\n")
	v.sb.Indent()
	methods, err := find.Methods(program)
	if err != nil {
		return errors.Wrap(err, "error finding methods")
	}

	for _, method := range methods {
		err := method.AcceptMethodVisitor(v)
		if err != nil {
			return errors.Wrap(err, "error serializing method")
		}
	}

	v.sb.Dedent()
	v.sb.WriteString("\n")

	v.sb.WriteString(":BODY\n")
	v.sb.Indent()
	err = program.Entrypoint().AcceptValueVisitor(v)
	v.sb.Dedent()
	if err != nil {
		return errors.Wrap(err, "failed to serialize statement")
	}
	return nil
}

func (v *visitor) serializeParameter(p graphite.Parameter) {
	v.sb.WriteString(fmt.Sprintf("{[%s] %s}", p.ReturnType().Name(), p.Name()))
}
