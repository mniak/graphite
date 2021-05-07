package graphite

type statement struct {
	method     *methodDeclaration
	args       []Argument
	returnType Type
}

func (s statement) ReturnType() Type {
	return s.returnType
}

func MethodInvocation(method *methodDeclaration, args []Argument) statement {
	return statement{
		method:     method,
		args:       args,
		returnType: method.ReturnType(),
	}
}

func BinaryOperation(operator string, left IValue, right IValue) statement {
	return statement{}
}
