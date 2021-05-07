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

func TypeString() Type {
	return primitiveType{"String"}
}

func TypeInt32() Type {
	return primitiveType{"Int32"}
}
