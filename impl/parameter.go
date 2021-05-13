package impl

import "github.com/mniak/graphite"

type parameter struct {
	name    string
	theType graphite.Type
}

func (p *parameter) Name() string {
	return p.name
}

func (p *parameter) Type() graphite.Type {
	return p.theType
}

func NewParameter(name string, theType graphite.Type) *parameter {
	return &parameter{
		name:    name,
		theType: theType,
	}
}
