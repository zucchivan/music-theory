// Chords are expressed in readable strings, e.g. CMb5b7 or Cm679-5
package chord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToYAML(t *testing.T) {
	c := Of("Cm769-5")
	out := c.ToYAML()
	assert.Equal(t, "root: C\ntones:\n  1: C\n  3: D#\n  6: A\n  7: A#\n  9: D\n", out)
}
