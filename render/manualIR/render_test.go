package manualIR

import (
	"github.com/mniak/graphite/render/samples"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleProgram(t *testing.T) {
	code, err := SerializeProgram(samples.SimpleProgram())
	assert.NoError(t, err)
	assert.Equal(t, `define i32 @f(i32 %param_a, i32 %param_b) {
  %var_1 = mul i32 2, %param_b
  %var_0 = add i32 %param_a, %var_1
  ret i32 %var_0
}

define i32 @main() {
  %var_0 = call i32 @f(i32 10, i32 20)
  ret i32 %var_0
}
`, code)
}
