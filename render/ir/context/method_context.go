package context

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
)

type MethodContext interface {
	ProgramContext
	RegisterParameter(param graphite.Parameter, value *llvmir.Param) *llvmir.Param
	GetParameter(param graphite.Parameter) value.Value
}

type methodContext struct {
	ProgramContext
	paramMap map[graphite.Parameter]*llvmir.Param
}

func NewMethodContext(ctx ProgramContext) MethodContext {
	return &methodContext{
		ProgramContext: ctx,
		paramMap:       make(map[graphite.Parameter]*llvmir.Param, 0),
	}
}

func (c *methodContext) RegisterParameter(param graphite.Parameter, value *llvmir.Param) *llvmir.Param {
	c.paramMap[param] = value
	return value
}

func (c *methodContext) GetParameter(param graphite.Parameter) value.Value {
	return c.paramMap[param]
}
