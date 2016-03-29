// Keys are expressed in readable strings, e.g. C major or Ab minor
package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToYAML(t *testing.T) {
	c := Of("C major")
	out := c.ToYAML()
	assert.Equal(t, "root: C\nmode: Major\n", out)
}
