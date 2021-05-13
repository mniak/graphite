package ir

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/render/ir/context"
)

type methodVisitor struct {
	irModule *llvmir.Module
	irMethod *llvmir.Func
	scope    scope
	context  context.ProgramContext
}

func newMethodVisitor(m *llvmir.Module, context context.ProgramContext) graphite.MethodVisitor {
	return methodVisitor{
		irModule: m,
		scope:    newScope(),
		context:  context,
	}
}

func (v methodVisitor) VisitInternalMethod(m graphite.InternalMethod) error {
	irType, err := getIrType(m.Type())
	if err != nil {
		return err
	}
	params := m.Parameters()
	irParams := make([]*llvmir.Param, len(params))
	for i, param := range params {
		irType, err = getIrType(param.Type())
		if err != nil {
			return err
		}
		irParam := llvmir.NewParam(param.Name(), irType)
		irParams[i] = irParam
	}
	irFn := v.irModule.NewFunc(m.Name(), irType, irParams...)
	irBody := irFn.NewBlock("body")
	vv := newValueVisitor(irBody, v.scope, v.context.NewMethodContext())
	err = m.Body().AcceptValueVisitor(vv)
	if err != nil {
		return err
	}
	irBody.NewRet(vv.lastIrValue)
	v.irMethod = v.context.RegisterFunction(m, irFn)
	return nil
}

func (v methodVisitor) VisitNativeOperation(m graphite.NativeOperation) error {
	return nil
}
