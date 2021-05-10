package graphite

type Value interface {
	Dispatcher
	ReturnType() Type
}
