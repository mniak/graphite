package main

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/demo/serialization"
	"github.com/mniak/graphite/impl"
	"github.com/mniak/graphite/native"
	"log"
)

func main() {
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

	code, err := serialization.Serialize(program)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(code)
}
