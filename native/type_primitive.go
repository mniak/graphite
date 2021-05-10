package native

type primitiveType struct {
	name string
}

func (p primitiveType) IsPrimitive() bool {
	return true
}

func (p primitiveType) Name() string {
	return p.name
}

func TypeString() primitiveType {
	return primitiveType{"String"}
}

func TypeInt32() primitiveType {
	return primitiveType{"Int32"}
}
