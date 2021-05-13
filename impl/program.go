package impl

import (
	"github.com/mniak/graphite"
)

type program struct {
	methods    []graphite.Method
	entrypoint graphite.Value
}

func (p program) Entrypoint() graphite.Value {
	return p.entrypoint
}

func ProgramWithoutLibraries(methods []graphite.Method, entrypoint graphite.Value) program {
	return program{
		methods:    methods,
		entrypoint: entrypoint,
	}
}
