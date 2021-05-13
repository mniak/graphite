package ir

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/render/ir/context"
	"github.com/mniak/graphite/render/ir/wrappers"
)

type methodVisitor struct {
	irModule *llvmir.Module
	irMethod *llvmir.Func
	scope    scope
	context  context.ProgramContext
}

func newMethodVisitor(m *llvmir.Module, context context.ProgramContext) wrappers.IRMethodVisitor {
	return methodVisitor{
		irModule: m,
		scope:    newScope(),
		context:  context,
	}
}

func (v methodVisitor) VisitInternalMethod(m graphite.InternalMethod) (value.Value, error) {
	irType, err := getIrType(m.Type())
	if err != nil {
		return nil, err
	}
	params := m.Parameters()
	irParams := make([]*llvmir.Param, len(params))
	for i, param := range params {
		irType, err = getIrType(param.Type())
		if err != nil {
			return nil, err
		}
		irParam := llvmir.NewParam(param.Name(), irType)
		irParams[i] = irParam
	}
	irFn := v.irModule.NewFunc(m.Name(), irType, irParams...)
	irBody := irFn.NewBlock("body")
	vv := newValueVisitor(irBody, v.scope, v.context.NewMethodContext())
	val, err := wrappers.WrapValueDispatcher(m.Body()).AcceptValueVisitor(vv)
	if err != nil {
		return nil, err
	}
	irBody.NewRet(val)
	return v.context.RegisterFunction(m, irFn), nil
}

func (v methodVisitor) VisitNativeOperation(m graphite.NativeOperation) (value.Value, error) {
	return nil, nil
}
