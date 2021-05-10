package graphite

type Dispatcher interface {
	Accept(visitor Visitor) error
}

type Visitor interface {
	VisitProgram(p Program) error
	VisitMethodInvocation(mi Invocation) error
	VisitInt32Literal(i int32) error
	VisitParameterValue(pv ParameterValue) error
	VisitArgument(a Argument) error
}
