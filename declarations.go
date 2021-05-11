package graphite

type Program interface {
	Entrypoint() Value
}

type Parameter interface {
	Name() string
	ReturnType() Type
}
