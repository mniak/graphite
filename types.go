package graphite

type Type interface {
	IsPrimitive() bool
	GetName() string
}

type primitiveType struct {
	name string
}

func (p primitiveType) IsPrimitive() bool {
	return true
}

func (p primitiveType) GetName() string {
	return p.name
}

func PrimitiveString() Type {
	return primitiveType{"String"}
}
