package ir

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/mniak/graphite"
)

type methodVisitor struct {
	irModule *llvmir.Module
	irMethod *llvmir.Func
	scope    scope
}

func newMethodVisitor(m *llvmir.Module) graphite.MethodVisitor {
	return methodVisitor{
		irModule: m,
		scope:    newScope(),
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
	vv := newValueVisitor(irBody, v.scope)
	err = m.Body().AcceptValueVisitor(vv)
	if err != nil {
		return err
	}
	irBody.NewRet(vv.lastIrValue)
	v.irMethod = irFn
	return nil
}

func (v methodVisitor) VisitNativeOperation(m graphite.NativeOperation) error {
	return nil
}
