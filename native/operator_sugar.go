package native

import (
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/impl"
)

func Int32Add(a graphite.Value, b graphite.Value) graphite.Invocation {
	operator := OperatorInt32Addition()
	params := operator.Parameters()
	return impl.NewInvocation(
		operator,
		[]graphite.Argument{
			impl.NewArgument(params[0], a),
			impl.NewArgument(params[1], b),
		},
	)
}

func Int32Mult(a graphite.Value, b graphite.Value) graphite.Invocation {
	operator := OperatorInt32Multiplication()
	params := operator.Parameters()
	return impl.NewInvocation(
		operator,
		[]graphite.Argument{
			impl.NewArgument(params[0], a),
			impl.NewArgument(params[1], b),
		},
	)
}
