package ir

import (
	"fmt"
	llvmir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/native"
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
	argsMapByParam := make(map[graphite.Parameter]value.Value)
	argsMapByName := make(map[string]value.Value)
	for _, arg := range args {
		irArg, err := wrappers.WrapValueDispatcher(arg.Value()).AcceptValueVisitor(v)
		if err != nil {
			return nil, err
		}
		argsMapByParam[arg.Parameter()] = irArg
		argsMapByName[arg.Parameter().Name()] = irArg
	}
	method := i.Method()
	if method.IsNative() {
		methodName := method.Name()
		switch methodName {
		case "*":
			return v.irBlock.NewMul(argsMapByName[native.BINARY_OPERATOR_FIRST_PARAM], argsMapByName[native.BINARY_OPERATOR_SECOND_PARAM]), nil
		case "+":
			return v.irBlock.NewAdd(argsMapByName[native.BINARY_OPERATOR_FIRST_PARAM], argsMapByName[native.BINARY_OPERATOR_SECOND_PARAM]), nil
		default:
			return nil, fmt.Errorf("could not produce invocation of native method %s", methodName)
		}
	} else {
		irFunc := v.context.GetFunction(i.Method())
		irArgs := make([]value.Value, len(argsMapByParam))
		for idx, param := range i.Method().Parameters() {
			irArgs[idx] = argsMapByParam[param]
		}
		return v.irBlock.NewCall(irFunc, irArgs...), nil
	}
}

func (v valueVisitor) VisitParameterValue(pv graphite.ParameterValue) (value.Value, error) {
	return v.context.GetParameter(pv.Parameter()), nil
}

func (v valueVisitor) VisitInt32Literal(i int32) (value.Value, error) {
	return constant.NewInt(types.I32, int64(i)), nil
}
