package find

import (
	"github.com/mniak/graphite"
)

type findVisitor struct {
	match   func(interface{}) bool
	results []interface{}
}

func (v *findVisitor) check(x interface{}) {
	if v.match(x) {
		v.results = append(v.results, x)
	}
}
func (v *findVisitor) VisitProgram(p graphite.Program) error {
	v.check(p)
	return p.Entrypoint().AcceptValueVisitor(v)
}

func (v *findVisitor) VisitInvocation(i graphite.Invocation) error {
	v.check(i)
	err := i.Method().AcceptMethodVisitor(v)
	if err != nil {
		return err
	}
	args := i.Arguments()
	for _, arg := range args {
		err := v.findInArgument(arg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *findVisitor) VisitParameterValue(pv graphite.ParameterValue) error {
	v.check(pv)
	v.check(pv.Parameter())
	return nil
}

func (v *findVisitor) VisitInt32Literal(i int32) error {
	v.check(i)
	return nil
}

func (v *findVisitor) VisitInternalMethod(m graphite.InternalMethod) error {
	v.check(m)
	for _, param := range m.Parameters() {
		v.check(param)
	}
	return nil
}

func (v *findVisitor) findInArgument(a graphite.Argument) error {
	v.check(a)
	v.check(a.Parameter())
	return a.Value().AcceptValueVisitor(v)
}
