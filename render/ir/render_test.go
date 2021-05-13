package ir

import (
	"github.com/mniak/graphite/render/samples"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleProgram(t *testing.T) {
	code, err := SerializeProgram(samples.SimpleProgram())
	assert.NoError(t, err)
	assert.Equal(t, `define i32 @f(i32 %a, i32 %b) {
body:
	%0 = mul i32 2, %b
	%1 = add i32 %a, %0
	ret i32 %1
}

define i32 @main() {
body:
	%0 = call i32 @f(i32 10, i32 20)
	ret i32 %0
}
`, code)
}
