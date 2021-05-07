package main

import (
	"fmt"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/demo/serialization"
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

	paramA := graphite.MethodParameterDeclaration{
		Name:    "a",
		TheType: graphite.TypeInt32(),
	}
	paramB := graphite.MethodParameterDeclaration{
		Name:    "b",
		TheType: graphite.TypeInt32(),
	}
	methodF := graphite.MethodDeclaration(
		"f",
		[]graphite.MethodParameterDeclaration{paramA, paramB},
		graphite.BinaryOperation(
			"+",
			graphite.ValueFromParameter(paramA),
			graphite.BinaryOperation(
				"x",
				2,
				graphite.ValueFromParameter(paramB),
			),
		),
	)

	entryPoint := methodF.Invocation([]graphite.Argument{
		{
			Parameter: paramA,
			Value:     graphite.Int32Literal{Value: 10},
		},
		{
			Parameter: paramB,
			Value:     graphite.Int32Literal{Value: 20},
		},
	})

	program := graphite.ProgramWithoutLibraries(entryPoint)

	code, err := serialization.SerializeProgram(program)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(code)
}
