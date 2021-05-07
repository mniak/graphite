package graphite

type ModuleDeclaration struct {
	name    string
	methods []MethodDeclaration
}
type MethodDeclaration struct {
	name       string
	parameters []MethodParameterDeclaration
	statement  IStatement
}

func (m MethodDeclaration) GetName() string {
	return m.name
}

type MethodParameterDeclaration struct {
	Name    string
	TheType Type
}

type ExternalLibraryDeclaration struct {
	Name    string
	Modules []ExternalModuleDeclaration
}
type ExternalModuleDeclaration struct {
	Name    string
	Methods []ExternalMethodDeclaration
}
type ExternalMethodDeclaration struct {
	Name       string
	Parameters []MethodParameterDeclaration
}

func (e ExternalMethodDeclaration) GetName() string {
	return e.Name
}

type IMethodDeclaration interface {
	GetName() string
}
