package lisp

import (
	"github.com/mniak/graphite/render/samples"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleProgram(t *testing.T) {
	code, err := SerializeProgram(samples.SimpleProgram())
	assert.NoError(t, err)
	assert.Equal(t, `(defun f (a b)
  (+ a (* 2 b)))

(f 10 20)
`, code)
}
