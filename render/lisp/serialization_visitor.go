package lisp

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/render/writer"
	"github.com/pkg/errors"
	"strconv"
)

type visitor struct {
	writer writer.Writer
}

func (v *visitor) VisitNativeOperation(graphite.NativeOperation) error {
	return nil
}

func (v *visitor) String() string {
	return v.writer.String()
}

func (v *visitor) VisitInternalMethod(m graphite.InternalMethod) error {
	v.writer.WriteString(fmt.Sprintf("(defun %s (", m.Name()))
	params := m.Parameters()
	first := true
	for _, param := range params {
		if first {
			first = false
		} else {
			v.writer.WriteString(" ")
		}
		v.writer.WriteString(param.Name())
	}
	v.writer.WriteString(")\n")
	v.writer.Indent()

	err := m.Body().AcceptValueVisitor(v)

	v.writer.Dedent()
	v.writer.WriteString(")\n")

	return err
}

func (v *visitor) VisitInvocation(mi graphite.Invocation) error {
	args := mi.Arguments()
	v.writer.WriteString("(")
	v.writer.WriteString(mi.Method().Name())
	for _, arg := range args {
		v.writer.WriteString(" ")
		err := v.serializeArgument(arg)
		if err != nil {
			return errors.Wrap(err, "failed to serialize argument")
		}
	}
	v.writer.WriteString(")")
	return nil
}

func (v *visitor) VisitInt32Literal(i int32) error {
	v.writer.WriteString(strconv.Itoa(int(i)))
	return nil
}

func (v *visitor) VisitParameterValue(pv graphite.ParameterValue) error {
	v.writer.WriteString(pv.Parameter().Name())
	return nil
}

func (v *visitor) serializeArgument(a graphite.Argument) error {
	return a.Value().AcceptValueVisitor(v)
}
