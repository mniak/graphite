package context

import (
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
)

type MethodContext interface {
	ProgramContext
	RegisterParameter(param graphite.Parameter, value value.Value) value.Value
	GetParameter(param graphite.Parameter) value.Value
}

type methodContext struct {
	ProgramContext
	paramMap map[graphite.Parameter]value.Value
}

func NewMethodContext(ctx ProgramContext) MethodContext {
	return &methodContext{
		ProgramContext: ctx,
		paramMap:       make(map[graphite.Parameter]value.Value, 0),
	}
}

func (c *methodContext) RegisterParameter(param graphite.Parameter, value value.Value) value.Value {
	c.paramMap[param] = value
	return value
}

func (c *methodContext) GetParameter(param graphite.Parameter) value.Value {
	return c.paramMap[param]
}
