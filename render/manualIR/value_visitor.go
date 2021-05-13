package manualIR

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mniak/graphite"
	"github.com/mniak/graphite/render/writer"
	"github.com/pkg/errors"
)

type valueVisitor struct {
	writer         writer.Writer
	varcount       int
	lastExpression string
}

func (v *valueVisitor) WriteString(str string) {
	v.writer.WriteString(str)
}

func (v *valueVisitor) Indent() {
	v.writer.Indent()
}

func (v *valueVisitor) Dedent() {
	v.writer.Dedent()
}

func (v *valueVisitor) newvar() string {
	varname := formatVariableName(strconv.Itoa(v.varcount))
	v.varcount++
	return varname
}

func (v *valueVisitor) VisitInvocation(mi graphite.Invocation) error {
	varname := v.newvar()
	method := mi.Method()
	native := method.IsNative()

	args := mi.Arguments()
	args2 := make([]string, len(args))
	for i, arg := range args {
		value := arg.Value()
		err := value.AcceptValueVisitor(v)
		if err != nil {
			return errors.Wrap(err, "failed to serialize argument")
		}
		irType, err := getIrType(value.ReturnType())
		if err != nil {
			return errors.Wrap(err, "failed to serialize argument type")
		}
		if native {
			args2[i] = v.lastExpression
		} else {
			args2[i] = fmt.Sprintf("%s %s", irType, v.lastExpression)
		}
	}

	arglist := strings.Join(args2, ", ")

	if native {
		instr, err := getInstructionName(method)
		if err != nil {
			return errors.Wrap(err, "failed to serialize native instruction name")
		}
		v.WriteString(fmt.Sprintf("%s = %s %s\n", varname, instr, arglist))
	} else {
		irType, err := getIrType(mi.ReturnType())
		if err != nil {
			return errors.Wrap(err, "error serializing invocation return type")
		}
		v.WriteString(fmt.Sprintf("%s = call %s %s(%s)\n", varname, irType, formatFunctionName(mi.Method().Name()), arglist))
	}
	v.lastExpression = varname
	return nil
}

func (v *valueVisitor) VisitParameterValue(pv graphite.ParameterValue) error {
	v.lastExpression = formatParameterName(pv.Parameter().Name())
	return nil
}

func (v *valueVisitor) VisitInt32Literal(i int32) error {
	v.lastExpression = strconv.Itoa(int(i))
	return nil
}
