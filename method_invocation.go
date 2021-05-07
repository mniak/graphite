package graphite

type MethodInvocation struct {
	Method    IMethodDeclaration
	Arguments []Argument
}
