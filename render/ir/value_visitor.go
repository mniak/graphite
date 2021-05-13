package ir

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/render/ir/context"
)

type valueVisitor struct {
	irBlock     *llvmir.Block
	scope       scope
	context     context.MethodContext
	lastIrValue value.Value
}

func newValueVisitor(b *llvmir.Block, s scope, context context.MethodContext) valueVisitor {
	return valueVisitor{
		irBlock: b,
		scope:   s,
		context: context,
	}
}

func (v valueVisitor) VisitInvocation(i graphite.Invocation) error {
	args := i.Arguments()
	irArgs := make([]value.Value, len(args))
	for idx, arg := range args {
		err := arg.Value().AcceptValueVisitor(v)
		if err != nil {
			return err
		}
		irArgs[idx] = v.lastIrValue
	}
	irFunc := v.context.GetFunction(i.Method())
	v.irBlock.NewCall(irFunc, irArgs...)
	return nil
}

func (v valueVisitor) VisitParameterValue(pv graphite.ParameterValue) error {
	param := pv.Parameter()
	v.lastIrValue = v.context.GetParameter(param)
	return nil
}

func (v valueVisitor) VisitInt32Literal(i int32) error {
	v.lastIrValue = constant.NewInt(types.I32, int64(i))
	return nil
}
