// Keys are expressed in readable strings, e.g. C major or Ab minor
package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToYAML(t *testing.T) {
	testKeySpecYAML(t, "C major", "root: C\nmode: Major\nrelative:\n  root: A\n  mode: Minor\n")
	testKeySpecYAML(t, "A minor", "root: A\nmode: Minor\nrelative:\n  root: C\n  mode: Major\n")
}

/*
 *
 private */

func testKeySpecYAML(t *testing.T, name string, expectOut string) {
	c := Of(name)
	out := c.ToYAML()
	assert.Equal(t, expectOut, out)
}
