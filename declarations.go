package graphite

type ModuleDeclaration struct {
	name    string
	methods []methodDeclaration
}

func MethodDeclaration(name string, parameters []MethodParameterDeclaration, statement statement) methodDeclaration {
	return methodDeclaration{
		name:       name,
		parameters: parameters,
		statement:  statement,
	}
}
func (m methodDeclaration) GetName() string {
	return m.name
}

func (m *methodDeclaration) Invocation(args []Argument) statement {
	return MethodInvocation(m, args)
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

type methodDeclaration struct {
	name       string
	parameters []MethodParameterDeclaration
	statement  statement
}

func (m methodDeclaration) ReturnType() Type {
	return m.statement.ReturnType()
}
