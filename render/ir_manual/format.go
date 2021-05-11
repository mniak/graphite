package ir_manual

import "fmt"

func formatVariableName(name string) string {
	return fmt.Sprintf("%%var_%s", name)
}

func formatParameterName(name string) string {
	return fmt.Sprintf("%%param_%s", name)
}

func formatFunctionName(name string) string {
	return fmt.Sprintf("@%s", name)
}
