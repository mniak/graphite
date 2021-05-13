package ir

import (
	llvmir "github.com/llir/llvm/ir"
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/find"
	"github.com/mniak/graphite/render/ir/context"
	"github.com/mniak/graphite/render/ir/wrappers"
	"github.com/pkg/errors"
)

func SerializeProgram(program graphite.Program) (string, error) {
	m := llvmir.NewModule()
	methods, err := find.Methods(program)
	if err != nil {
		return "", errors.Wrap(err, "error finding methods")
	}

	ctx := context.NewProgramContext()
	for _, method := range methods {
		methodVisitor := newMethodVisitor(m, ctx)
		_, err := wrappers.WrapMethodDispatcher(method).AcceptMethodVisitor(methodVisitor)
		if err != nil {
			return "", errors.Wrap(err, "error serializing method")
		}
	}

	//w.WriteString("\ndefine i32 @main() {\n")
	//w.Indent()
	//valueVisitor := newValueVisitor(w)
	//err = program.Entrypoint().AcceptValueVisitor(valueVisitor)
	//w.WriteString(fmt.Sprintf("ret i32 %s\n", valueVisitor.lastExpression))
	//if err != nil {
	//	return "", errors.Wrap(err, "failed to serialize statement")
	//}
	//
	//w.Dedent()
	//w.WriteString("}\n")

	return m.String(), nil
}
