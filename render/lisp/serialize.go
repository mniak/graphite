package lisp

import (
	"github.com/mniak/graphite"
	"github.com/mniak/graphite/find"
	"github.com/mniak/graphite/render/writer"
	"github.com/pkg/errors"
)

func SerializeProgram(program graphite.Program) (string, error) {
	w := writer.New()
	methods, err := find.Methods(program)
	if err != nil {
		return "", errors.Wrap(err, "error finding methods")
	}

	visitor := visitor{
		writer: w,
	}
	for _, method := range methods {

		err := method.AcceptMethodVisitor(&visitor)
		if err != nil {
			return "", errors.Wrap(err, "error serializing method")
		}
	}

	w.WriteString("\n")

	err = program.Entrypoint().AcceptValueVisitor(&visitor)
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize statement")
	}

	w.WriteString("\n")
	return w.String(), nil
}
