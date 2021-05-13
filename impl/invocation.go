package impl

import "github.com/mniak/graphite"

type invocation struct {
	method graphite.Method
	args   []graphite.Argument
}

func (i invocation) Arguments() []graphite.Argument {
	return i.args
}

func (i invocation) Method() graphite.Method {
	return i.method
}

func (i invocation) AcceptValueVisitor(visitor graphite.ValueVisitor) error {
	return visitor.VisitInvocation(i)
}

func (i invocation) ReturnType() graphite.Type {
	return i.method.Type()
}

func NewInvocation(method graphite.Method, args []graphite.Argument) invocation {
	return invocation{
		method: method,
		args:   args,
	}
}
