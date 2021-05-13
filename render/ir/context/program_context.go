package context

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/mniak/graphite"
)

type ProgramContext interface {
	RegisterFunction(method graphite.Method, fn *llvmir.Func) *llvmir.Func
	GetFunction(method graphite.Method) *llvmir.Func
	NewMethodContext() MethodContext
}

type programContext struct {
	funcMap map[graphite.Method]*llvmir.Func
}

func NewProgramContext() ProgramContext {
	return &programContext{
		funcMap: make(map[graphite.Method]*llvmir.Func, 0),
	}
}

func (c *programContext) RegisterFunction(method graphite.Method, fn *llvmir.Func) *llvmir.Func {
	c.funcMap[method] = fn
	return fn
}

func (c *programContext) GetFunction(method graphite.Method) *llvmir.Func {
	return c.funcMap[method]
}

func (c *programContext) NewMethodContext() MethodContext {
	return NewMethodContext(c)
}
