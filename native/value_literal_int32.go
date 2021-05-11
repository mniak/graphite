package native

import (
	"github.com/mniak/graphite"
)

type int32Literal struct {
	value int32
}

func (i int32Literal) AcceptValueVisitor(visitor graphite.ValueVisitor) error {
	return visitor.VisitInt32Literal(i.value)
}

func (i int32Literal) ReturnType() graphite.Type {
	return TypeInt32()
}

func Int32(value int32) int32Literal {
	return int32Literal{value: value}
}
