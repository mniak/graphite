package serialization

import (
	"fmt"
	"github.com/mniak/graphite"
	"strings"
)

func SerializeProgram(program graphite.Program) (string, error) {
	var err error
	statements := make([]string, len(program.Entrypoint))
	for i, s := range program.Entrypoint {
		statements[i], err = serializeStatement(s)
		if err != nil {
			return "", err
		}
	}
	return strings.Join(statements, "\n"), nil
}

func serializeStatement(s graphite.IStatement) (string, error) {
	switch st := s.(type) {
	case graphite.MethodInvocation:
		return serializeMethodInvocation(st)
	default:
		return "", fmt.Errorf("statement could not be serialized: %s", s)
	}
}

func serializeMethodInvocation(mi graphite.MethodInvocation) (string, error) {
	var sb strings.Builder
	sb.WriteString(mi.Method.GetName())
	sb.WriteString("{")
	for _, arg := range mi.Arguments {
		sb.WriteString(arg.Parameter.Name)
		sb.WriteString("=")
		value, err := SerializeValue(arg.Value)
		if err != nil {
			return "", err
		}
		sb.WriteString(value)
	}
	sb.WriteString("}")
	return sb.String(), nil
}

func SerializeValue(v graphite.IValue) (string, error) {
	switch vt := v.(type) {
	case graphite.StringLiteral:
		return vt.Value, nil
	default:
		return "", fmt.Errorf("value could not be serialized: %v", v)
	}
}
