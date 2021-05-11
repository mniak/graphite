package graphite

type MethodDispatcher interface {
	AcceptMethodVisitor(visitor MethodVisitor) error
}
type MethodVisitor interface {
	VisitInternalMethod(m InternalMethod) error
	VisitNativeOperator(o NativeOperator) error
}

type Method interface {
	MethodDispatcher
	Name() string
	Parameters() []Parameter
	ReturnType() Type
}

type InternalMethod interface {
	Method
}
type NativeOperator interface {
	Method
}
