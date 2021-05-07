package graphite

type program struct {
	libraries  []ExternalLibraryDeclaration
	entrypoint statement
}

func Program(libraries []ExternalLibraryDeclaration, entrypoint statement) program {
	return program{
		libraries:  libraries,
		entrypoint: entrypoint,
	}
}

func ProgramWithoutLibraries(entrypoint statement) program {
	return program{
		libraries:  []ExternalLibraryDeclaration{},
		entrypoint: entrypoint,
	}
}
