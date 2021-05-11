package render

import (
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/impl"
	"github.com/mniak/graphite/native"
	"github.com/mniak/graphite/render/ir_manual"
	"github.com/mniak/graphite/render/lisp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MakeProgram() graphite.Program {
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
	return program
}

func TestIR(t *testing.T) {
	program := MakeProgram()
	code, err := ir_manual.SerializeProgram(program)
	assert.NoError(t, err)
	assert.Equal(t, `define i32 @f(i32 %param_a, i32 %param_b) {
  %var_1 = mul i32 2, %param_b
  %var_0 = add i32 %param_a, %var_1
  ret i32 %var_0
}

define i32 @main() {
  %var_0 = call i32 @f(i32 10, i32 20)
  ret i32 %var_0
}
`, code)
}

func TestLisp(t *testing.T) {
	program := MakeProgram()
	code, err := lisp.SerializeProgram(program)
	assert.NoError(t, err)
	assert.Equal(t, `(defun f (a b)
  (+ a (* 2 b)))

(f 10 20)
`, code)
}
