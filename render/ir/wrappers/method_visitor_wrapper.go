package wrappers

import (
	"github.com/llir/llvm/ir/value"
	"github.com/mniak/graphite"
)

type IRMethodVisitor interface {
	VisitInternalMethod(graphite.InternalMethod) (value.Value, error)
	VisitNativeOperation(graphite.NativeOperation) (value.Value, error)
}

type irMethodVisitor struct {
	inner     IRMethodVisitor
	lastValue value.Value
}

func (w irMethodVisitor) VisitInternalMethod(m graphite.InternalMethod) error {
	val, err := w.inner.VisitInternalMethod(m)
	w.lastValue = val
	return err
}

func (w irMethodVisitor) VisitNativeOperation(o graphite.NativeOperation) error {
	val, err := w.inner.VisitNativeOperation(o)
	w.lastValue = val
	return err
}

func WrapMethodVisitor(v IRMethodVisitor) irMethodVisitor {
	return irMethodVisitor{
		inner: v,
	}
}

type IRMethodDispatcher interface {
	AcceptMethodVisitor(IRMethodVisitor) (value.Value, error)
}

type irMethodDispatcher struct {
	inner graphite.MethodDispatcher
}

func (d irMethodDispatcher) AcceptMethodVisitor(v IRMethodVisitor) (value.Value, error) {
	wrapper := WrapMethodVisitor(v)
	err := d.inner.AcceptMethodVisitor(wrapper)
	return wrapper.lastValue, err
}

func WrapMethodDispatcher(d graphite.MethodDispatcher) IRMethodDispatcher {
	return irMethodDispatcher{
		inner: d,
	}
}
