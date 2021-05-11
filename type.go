package graphite

type Type interface {
	Name() string
	IsPrimitive() bool
}
