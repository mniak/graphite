package serialization

import (
	"fmt"
	"github.com/mniak/graphite"
	"strings"
)

func SerializeProgram(program graphite.Program) (string, error) {
	return serializeStatement(program.entrypoin)
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
		value, err := serializeValue(arg.Value)
		if err != nil {
			return "", err
		}
		sb.WriteString(value)
	}
	sb.WriteString("}")
	return sb.String(), nil
}

func serializeValue(v graphite.IValue) (string, error) {
	switch vt := v.(type) {
	case graphite.StringLiteral:
		return vt.Value, nil
	default:
		return "", fmt.Errorf("value could not be serialized: %v", v)
	}
}
