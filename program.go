package graphite

type Program struct {
	Libraries  []ExternalLibraryDeclaration
	Entrypoint []IStatement
}
