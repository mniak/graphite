package manIR

import (
	"github.com/mniak/graphite"
	"testing"
)

func TestInterface(t *testing.T) {
	var theType programVisitor
	var theInteface graphite.AllVisitors

	theInteface = &theType
	_ = theInteface
}
