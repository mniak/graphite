package ir

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
)

type valueVisitor struct {
	irBlock     *llvmir.Block
	scope       scope
	varmap      map[string]string
	lastIrValue value.Value
}

func newValueVisitor(b *llvmir.Block, s scope, varmap map[string]string) valueVisitor {
	return valueVisitor{
		irBlock: b,
		scope:   s,
	}
}

func (v valueVisitor) VisitInvocation(i graphite.Invocation) error {
	// v.irBlock.NewCall()
	return nil
}

func (v valueVisitor) VisitParameterValue(pv graphite.ParameterValue) error {
	param := pv.Parameter()
	paramName := v.varmap[param.Name()]
	return nil
}

func (v valueVisitor) VisitInt32Literal(i int32) error {
	v.lastIrValue = constant.NewInt(types.I32, int64(i))
	return nil
}
