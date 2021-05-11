package lisp

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/find"
	"github.com/pkg/errors"
	"strconv"
)

type visitor struct {
	sb indentedStringBuilder
}

func (v *visitor) String() string {
	return v.sb.String()
}

func (v *visitor) VisitInternalMethod(m graphite.InternalMethod) error {
	v.sb.WriteString(fmt.Sprintf("(defun %s (", m.Name()))
	params := m.Parameters()
	first := true
	for _, param := range params {
		if first {
			first = false
		} else {
			v.sb.WriteString(" ")
		}
		v.sb.WriteString(param.Name())
	}
	v.sb.WriteString(")\n")
	v.sb.Indent()

	err := m.Body().AcceptValueVisitor(v)

	v.sb.Dedent()
	v.sb.WriteString(")\n")

	return err
}

func (v *visitor) VisitInvocation(mi graphite.Invocation) error {
	args := mi.Arguments()
	v.sb.WriteString("(")
	v.sb.WriteString(mi.Method().Name())
	for _, arg := range args {
		v.sb.WriteString(" ")
		err := v.serializeArgument(arg)
		if err != nil {
			return errors.Wrap(err, "failed to serialize argument")
		}
	}
	v.sb.WriteString(")")
	return nil
}

func (v *visitor) VisitInt32Literal(i int32) error {
	v.sb.WriteString(strconv.Itoa(int(i)))
	return nil
}

func (v *visitor) VisitParameterValue(pv graphite.ParameterValue) error {
	v.sb.WriteString(pv.Parameter().Name())
	return nil
}

func (v *visitor) serializeArgument(a graphite.Argument) error {
	return a.Value().AcceptValueVisitor(v)
}
func (v *visitor) serializeProgram(program graphite.Program) error {
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

	v.sb.WriteString("\n")

	err = program.Entrypoint().AcceptValueVisitor(v)
	if err != nil {
		return errors.Wrap(err, "failed to serialize statement")
	}
	return nil
}
