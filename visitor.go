package graphite

type AllVisitors interface {
	MethodVisitor
	ValueVisitor
}
