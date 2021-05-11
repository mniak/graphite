package writer

import (
	"fmt"
	"strings"
)

type Writer interface {
	fmt.Stringer
	WriteString(string)
	Indent()
	Dedent()
}

type writer struct {
	sb               *strings.Builder
	indentation      int
	endedWithNewLine bool
}

func New() Writer {
	return &writer{
		sb: new(strings.Builder),
	}
}

func (w *writer) WriteString(str string) {
	lines := strings.Split(str, "\n")
	spaces := strings.Repeat(" ", w.indentation)
	endnl := false
	for i, line := range lines {
		isfirst := i == 0
		if isfirst && w.endedWithNewLine {
			w.sb.WriteString(spaces)
		}

		islast := i == len(lines)-1
		if !islast {
			w.sb.WriteString(line)
			w.sb.WriteString("\n")
		} else {
			if line == "" {
				endnl = true
			} else {
				w.sb.WriteString(line)
			}
		}
	}
	w.endedWithNewLine = endnl
}

const INDENTATION = 2

func (w *writer) Indent() {
	w.indentation += INDENTATION
}

func (w *writer) Dedent() {
	if w.indentation >= INDENTATION {
		w.indentation -= INDENTATION
	}
}
func (w *writer) String() string {
	return w.sb.String()
}
