package native

import (
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/impl"
)

type binaryOperation struct {
	name   string
	aType  graphite.Type
	bType  graphite.Type
	native bool
}

func (o *binaryOperation) IsNative() bool {
	return o.native
}

func (o *binaryOperation) AcceptMethodVisitor(visitor graphite.MethodVisitor) error {
	return visitor.VisitNativeOperation(o)
}

const (
	BINARY_OPERATOR_FIRST_PARAM  = "a"
	BINARY_OPERATOR_SECOND_PARAM = "b"
)

func (o *binaryOperation) Parameters() []graphite.Parameter {
	return []graphite.Parameter{
		impl.NewParameter(BINARY_OPERATOR_FIRST_PARAM, TypeInt32()),
		impl.NewParameter(BINARY_OPERATOR_SECOND_PARAM, TypeInt32()),
	}
}

func (o *binaryOperation) Name() string {
	return o.name
}

func (o *binaryOperation) Type() graphite.Type {
	return TypeInt32()
}

var int32Mult = binaryOperation{
	name:   "*",
	aType:  TypeInt32(),
	bType:  TypeInt32(),
	native: true,
}

var int32Add = binaryOperation{
	name:   "+",
	aType:  TypeInt32(),
	bType:  TypeInt32(),
	native: true,
}

func OperatorInt32Addition() *binaryOperation {
	return &int32Add
}

func OperatorInt32Multiplication() *binaryOperation {
	return &int32Mult
}
