// It's possible to export a list of all known scale parsing rules.
package scale

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestListToYAML(t *testing.T) {
	c := ScaleModeList
	out := c.ToYAML()
	assert.Equal(t, "- Default (Major)\n- Minor\n- Major\n- Natural Minor\n- Diminished\n- Augmented\n- Melodic Minor Ascend\n- Melodic Minor Descend\n- Harmonic Minor\n- Ionian\n- Dorian\n- Phrygian\n- Lydian\n- Mixolydian\n- Aeolian\n- Locrian\n", out)
}
