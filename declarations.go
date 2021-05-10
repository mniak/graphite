package graphite

type Program interface {
	Dispatcher
	Entrypoint() Value
}

type Method interface {
	Name() string
	Parameters() []Parameter
	ReturnType() Type
}

type Parameter interface {
	Name() string
	ReturnType() Type
}

type Type interface {
	Name() string
}
