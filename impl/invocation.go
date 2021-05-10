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

func (i invocation) Accept(visitor graphite.Visitor) error {
	return visitor.VisitMethodInvocation(i)
}

func (i invocation) ReturnType() graphite.Type {
	return i.method.ReturnType()
}

func NewInvocation(method graphite.Method, args []graphite.Argument) invocation {
	return invocation{
		method: method,
		args:   args,
	}
}
