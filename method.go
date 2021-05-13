package graphite

type MethodDispatcher interface {
	AcceptMethodVisitor(visitor MethodVisitor) error
}

type MethodVisitor interface {
	VisitInternalMethod(m InternalMethod) error
	VisitNativeOperation(m NativeOperation) error
}

type Method interface {
	MethodDispatcher
	Name() string
	Parameters() []Parameter
	Type() Type
	IsNative() bool
}

type InternalMethod interface {
	Method
	Body() Value
}

type NativeOperation interface {
	Method
}
