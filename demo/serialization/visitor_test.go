package serialization

import (
	"github.com/mniak/graphite"
	"testing"
)

func TestInterface(t *testing.T) {
	var theType visitor
	var theInteface graphite.AllVisitors

	theInteface = &theType
	_ = theInteface
}
