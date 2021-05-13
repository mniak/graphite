package ir

import (
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/impl"
	"github.com/mniak/graphite/native"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleProgram(t *testing.T) {
	/*
		int f(int a, int b) {
			return a + 2*b;
		}

		int main() {
			return f(10, 20);
		}
	*/
	paramA := impl.NewParameter("a", native.TypeInt32())
	paramB := impl.NewParameter("b", native.TypeInt32())
	methodF := impl.NewInternalMethod(
		"f",
		[]graphite.Parameter{paramA, paramB},
		native.Int32Add(
			impl.ValueFromParameter(paramA),
			native.Int32Mult(
				native.Int32(2),
				impl.ValueFromParameter(paramB),
			),
		),
	)

	entryPoint := impl.NewInvocation(&methodF, []graphite.Argument{
		impl.NewArgument(paramA, native.Int32(10)),
		impl.NewArgument(paramB, native.Int32(20)),
	})

	program := impl.ProgramWithoutLibraries(entryPoint)

	code, err := SerializeProgram(program)
	assert.NoError(t, err)
	assert.Equal(t, `define i32 @f(i32 %a, i32 %b) {
body:
	%0 = mul i32 2, %b
	%1 = add i32 %a, %1
	ret i32 %0
}

define i32 @main() {
  %0 = call i32 @f(i32 10, i32 20)
  ret i32 %0
}
`, code)
}
