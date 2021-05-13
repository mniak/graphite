package ir

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/mniak/graphite"
)

func getIrType(p graphite.Type) (types.Type, error) {
	name := p.Name()
	if p.IsPrimitive() {
		switch name {
		case "Int32":
			return types.I32, nil
		default:
			return nil, fmt.Errorf("could not get IR equivalent for type %s", p.Name())
		}
	} else {
		return nil, fmt.Errorf("non primitive type %s cannot be converted to IR type. only primitives are supported at this moment", p.Name())
	}
}
