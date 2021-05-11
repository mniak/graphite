package graphite

type ValueDispatcher interface {
	AcceptValueVisitor(visitor ValueVisitor) error
}
type ValueVisitor interface {
	VisitInvocation(i Invocation) error
	VisitParameterValue(v ParameterValue) error
	VisitInt32Literal(i int32) error
}

type Value interface {
	ValueDispatcher
	ReturnType() Type
}
