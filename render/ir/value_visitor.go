package ir

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/render/ir/context"
	"github.com/mniak/graphite/render/ir/wrappers"
)

type valueVisitor struct {
	irBlock *llvmir.Block
	scope   scope
	context context.MethodContext
}

func newValueVisitor(b *llvmir.Block, s scope, context context.MethodContext) wrappers.IRValueVisitor {
	return valueVisitor{
		irBlock: b,
		scope:   s,
		context: context,
	}
}

func (v valueVisitor) VisitInvocation(i graphite.Invocation) (value.Value, error) {
	args := i.Arguments()
	irArgs := make([]value.Value, len(args))
	for idx, arg := range args {
		irArg, err := wrappers.WrapValueDispatcher(arg.Value()).AcceptValueVisitor(v)
		if err != nil {
			return nil, err
		}
		irArgs[idx] = irArg
	}
	irFunc := v.context.GetFunction(i.Method())

	return v.irBlock.NewCall(irFunc, irArgs...), nil
}

func (v valueVisitor) VisitParameterValue(pv graphite.ParameterValue) (value.Value, error) {
	return v.context.GetParameter(pv.Parameter()), nil
}

func (v valueVisitor) VisitInt32Literal(i int32) (value.Value, error) {
	return constant.NewInt(types.I32, int64(i)), nil
}
