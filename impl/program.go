package impl

import (
	"github.com/mniak/graphite"
)

type program struct {
	//libraries  []external
	entrypoint graphite.Value
}

func (p program) Entrypoint() graphite.Value {
	return p.entrypoint
}

//func Program(libraries []ExternalLibraryDeclaration, entrypoint graphite.Value) program {
//	return program{
//		libraries:  libraries,
//		entrypoint: entrypoint,
//	}
//}

func ProgramWithoutLibraries(entrypoint graphite.Value) program {
	return program{
		//libraries:  []ExternalLibraryDeclaration{},
		entrypoint: entrypoint,
	}
}
