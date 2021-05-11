package impl

import (
	"github.com/mniak/graphite"
	"testing"
)

func TestProgram_Interface(t *testing.T) {
	var theType program
	var theInterface graphite.Program

	theInterface = &theType
	_ = theInterface
}

func TestMethodInvocation_Interface(t *testing.T) {
	var theType invocation
	var theInterface graphite.Value

	theInterface = &theType
	_ = theInterface
}

func TestParameterValue_Interface(t *testing.T) {
	var theType parameterValue
	var theInterface graphite.ParameterValue

	theInterface = &theType
	_ = theInterface
}
