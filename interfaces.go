package graphite

type ParameterValue interface {
	Value
	Parameter() Parameter
}
