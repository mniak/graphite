package impl

import (
	"github.com/mniak/graphite"
)

type program struct {
	entrypoint graphite.Value
}

func (p program) Entrypoint() graphite.Value {
	return p.entrypoint
}

func ProgramWithoutLibraries(entrypoint graphite.Value) program {
	return program{
		entrypoint: entrypoint,
	}
}
