package native

import (
	"github.com/mniak/graphite"
	"testing"
)

func TestPrimitiveType_Interface(t *testing.T) {
	var theType primitiveType
	var theInterface graphite.Type

	theInterface = &theType
	_ = theInterface
}

func TestInt32Literal_Interface(t *testing.T) {
	var theType int32Literal
	var theInterface graphite.Value

	theInterface = &theType
	_ = theInterface
}
