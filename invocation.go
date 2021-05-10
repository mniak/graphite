package graphite

type Invocation interface {
	Value
	Method() Method
	Arguments() []Argument
}
type Argument interface {
	Dispatcher
	Parameter() Parameter
	Value() Value
}
