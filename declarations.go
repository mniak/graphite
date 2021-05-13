package graphite

type Program interface {
	Entrypoint() Value
}

type Parameter interface {
	Name() string
	Type() Type
}
