package impl

import "github.com/mniak/graphite"

type methodParameter struct {
	name    string
	theType graphite.Type
}

func (m methodParameter) Name() string {
	return m.name
}

func (m methodParameter) ReturnType() graphite.Type {
	return m.theType
}

func NewParameter(name string, theType graphite.Type) methodParameter {
	return methodParameter{
		name:    name,
		theType: theType,
	}
}
