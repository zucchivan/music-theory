// Chords are expressed in readable strings, e.g. CMb5b7 or Cm679-5
package chord

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestToYAML(t *testing.T) {
	c := Of("Cm769-5")
	out := c.ToYAML()
	assert.Equal(t, "root: C\ntones:\n  1: C\n  3: Eb\n  6: A\n  7: Bb\n  9: D\n", out)
}
