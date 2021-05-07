package graphite

type IValue interface {
}

type StringLiteral struct {
	Value string
}

type Int32Literal struct {
	Value int32
}

func ValueFromParameter(param MethodParameterDeclaration) IValue {
	return nil
}
