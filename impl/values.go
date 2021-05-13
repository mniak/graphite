package impl

import "github.com/mniak/graphite"

type parameterValue struct {
	param graphite.Parameter
}

func (p *parameterValue) Parameter() graphite.Parameter {
	return p.param
}

func (p *parameterValue) AcceptValueVisitor(visitor graphite.ValueVisitor) error {
	return visitor.VisitParameterValue(p)
}

func (p *parameterValue) ReturnType() graphite.Type {
	return p.param.Type()
}

func ValueFromParameter(param graphite.Parameter) graphite.ParameterValue {
	return &parameterValue{
		param: param,
	}
}
