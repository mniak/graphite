package graphite

type MethodDispatcher interface {
	AcceptMethodVisitor(visitor MethodVisitor) error
}
type MethodVisitor interface {
	VisitInternalMethod(m InternalMethod) error
}

type Method interface {
	MethodDispatcher
	Name() string
	Parameters() []Parameter
	ReturnType() Type
	IsNative() bool
}

type InternalMethod interface {
	Method
	Body() Value
}
type NativeOperator interface {
	Method
}
