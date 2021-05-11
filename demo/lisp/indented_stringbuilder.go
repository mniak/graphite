package lisp

import "strings"

type indentedStringBuilder struct {
	sb               strings.Builder
	indentation      int
	endedWithNewLine bool
}

func (isb *indentedStringBuilder) WriteString(str string) {
	lines := strings.Split(str, "\n")
	spaces := strings.Repeat(" ", isb.indentation)
	endnl := false
	for i, line := range lines {
		isfirst := i == 0
		if isfirst && isb.endedWithNewLine {
			isb.sb.WriteString(spaces)
		}

		islast := i == len(lines)-1
		if !islast {
			isb.sb.WriteString(line)
			isb.sb.WriteString("\n")
		} else {
			if line == "" {
				endnl = true
			} else {
				isb.sb.WriteString(line)
			}
		}
	}
	isb.endedWithNewLine = endnl
}

const INDENTATION = 2

func (isb *indentedStringBuilder) Indent() {
	isb.indentation += INDENTATION
}

func (isb *indentedStringBuilder) Dedent() {
	if isb.indentation >= INDENTATION {
		isb.indentation -= INDENTATION
	}
}
func (isb *indentedStringBuilder) String() string {
	return isb.sb.String()
}
