package lisp

import (
	"testing"

	"github.com/mniak/graphite"
)

func TestInterface(t *testing.T) {
	var theType visitor
	var theInteface graphite.AllVisitors

	theInteface = &theType
	_ = theInteface
}
