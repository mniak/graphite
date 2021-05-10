package serialization

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/pkg/errors"
)

type visitor struct {
	sb indentedStringBuilder
}

func (v *visitor) String() string {
	return v.sb.String()
}
func (v *visitor) VisitProgram(program graphite.Program) error {
	v.sb.WriteString(":DECLARATIONS\n")
	v.sb.Indent()
	v.sb.WriteString("not implemented")
	v.sb.Dedent()
	v.sb.WriteString("\n")

	v.sb.WriteString(":BODY\n")
	v.sb.Indent()
	err := program.Entrypoint().Accept(v)
	v.sb.Dedent()
	if err != nil {
		return errors.Wrap(err, "failed to serialize statement")
	}
	return nil
}

func (v *visitor) VisitMethodInvocation(mi graphite.Invocation) error {
	args := mi.Arguments()
	v.sb.WriteString(fmt.Sprintf("INVOKE %s (", mi.Method().Name()))
	for _, arg := range args {
		err := arg.Accept(v)
		if err != nil {
			return errors.Wrap(err, "failed to serialize argument")
		}
	}
	v.sb.WriteString(fmt.Sprintf(") -> %s;", mi.Method().ReturnType().Name()))
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

func (v *visitor) VisitArgument(a graphite.Argument) error {
	v.sb.WriteString(fmt.Sprintf("[arg] %s = ", a.Parameter().Name()))
	a.Value().Accept(v)
	return nil
}
