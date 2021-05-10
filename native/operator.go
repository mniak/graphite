package native

import (
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/impl"
)

type binaryOperator struct {
	name  string
	aType graphite.Type
	bType graphite.Type
}

func (o *binaryOperator) Parameters() []graphite.Parameter {
	return []graphite.Parameter{
		impl.NewParameter("a", TypeInt32()),
		impl.NewParameter("b", TypeInt32()),
	}
}

func (o *binaryOperator) Name() string {
	return o.name
}

func (o *binaryOperator) ReturnType() graphite.Type {
	return TypeInt32()
}

var int32Mult = binaryOperator{
	name:  "*",
	aType: TypeInt32(),
	bType: TypeInt32(),
}

var int32Add = binaryOperator{
	name:  "+",
	aType: TypeInt32(),
	bType: TypeInt32(),
}

func OperatorInt32Addition() graphite.Method {
	return &int32Add
}

func OperatorInt32Multiplication() graphite.Method {
	return &int32Mult
}
