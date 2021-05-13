package wrappers

import (
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
)

type IRValueVisitor interface {
	VisitInvocation(graphite.Invocation) (value.Value, error)
	VisitParameterValue(graphite.ParameterValue) (value.Value, error)
	VisitInt32Literal(int32) (value.Value, error)
}

type irValueVisitor struct {
	inner     IRValueVisitor
	lastValue value.Value
}

func (v irValueVisitor) VisitInvocation(i graphite.Invocation) error {
	val, err := v.inner.VisitInvocation(i)
	v.lastValue = val
	return err
}

func (v irValueVisitor) VisitParameterValue(pv graphite.ParameterValue) error {
	val, err := v.inner.VisitParameterValue(pv)
	v.lastValue = val
	return err
}

func (v irValueVisitor) VisitInt32Literal(i int32) error {
	val, err := v.inner.VisitInt32Literal(i)
	v.lastValue = val
	return err
}

func WrapValueVisitor(v IRValueVisitor) irValueVisitor {
	return irValueVisitor{
		inner: v,
	}
}

type IRValueDispatcher interface {
	AcceptValueVisitor(IRValueVisitor) (value.Value, error)
}

type irValueDispatcher struct {
	inner graphite.ValueDispatcher
}

func (d irValueDispatcher) AcceptValueVisitor(v IRValueVisitor) (value.Value, error) {
	wrapper := WrapValueVisitor(v)
	err := d.inner.AcceptValueVisitor(wrapper)
	return wrapper.lastValue, err
}

func WrapValueDispatcher(d graphite.ValueDispatcher) IRValueDispatcher {
	return irValueDispatcher{
		inner: d,
	}
}
