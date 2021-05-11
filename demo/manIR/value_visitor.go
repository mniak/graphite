package manIR

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type valueVisitor struct {
	parent         *programVisitor
	lastExpression string
	varcount       int
}

func (v *valueVisitor) WriteString(str string) {
	v.parent.sb.WriteString(str)
}

func (v *valueVisitor) Indent() {
	v.parent.sb.Indent()
}

func (v *valueVisitor) Dedent() {
	v.parent.sb.Dedent()
}

func (v *valueVisitor) newvar() string {
	varname := formatVariableName(strconv.Itoa(v.varcount))
	v.varcount++
	return varname
}

func (v *valueVisitor) VisitInvocation(mi graphite.Invocation) error {
	varname := v.newvar()

	args := mi.Arguments()
	args2 := make([]string, len(args))
	for i, arg := range args {
		arg.Value().AcceptValueVisitor(v)
		args2[i] = v.lastExpression
	}

	method := mi.Method()
	arglist := strings.Join(args2, " ")

	if method.IsNative() {
		instr, err := getInstructionName(method)
		if err != nil {
			return errors.Wrap(err, "failed to serialize native instruction name")
		}
		v.WriteString(fmt.Sprintf("%s = %s %s\n", varname, instr, arglist))
	} else {
		v.WriteString(fmt.Sprintf("%s = call %s (%s)\n", varname, mi.Method().Name(), arglist))
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
