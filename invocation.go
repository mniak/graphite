package graphite

type Invocation interface {
	Value
	Method() Method
	Arguments() []Argument
}

type Argument interface {
	Parameter() Parameter
	Value() Value
}
