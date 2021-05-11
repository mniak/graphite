package impl

import "github.com/mniak/graphite"

type internalMethod struct {
	name       string
	parameters []graphite.Parameter
	statement  graphite.Value
}

func (m *internalMethod) IsNative() bool {
	return false
}

func (m *internalMethod) Body() graphite.Value {
	return m.statement
}

func (m *internalMethod) AcceptMethodVisitor(visitor graphite.MethodVisitor) error {
	return visitor.VisitInternalMethod(m)
}

func NewInternalMethod(name string, parameters []graphite.Parameter, body graphite.Value) internalMethod {
	return internalMethod{
		name:       name,
		parameters: parameters,
		statement:  body,
	}
}

func (m *internalMethod) Name() string {
	return m.name
}

func (m *internalMethod) Parameters() []graphite.Parameter {
	return m.parameters
}

func (m *internalMethod) ReturnType() graphite.Type {
	return m.statement.ReturnType()
}
