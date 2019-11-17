// Key Difference to another Key can be calculated in semitones
package key

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestClass_Diff(t *testing.T) {
	assert.Equal(t, 5, Of("C#").Diff(Of("F#")))
	assert.Equal(t, -5, Of("F#").Diff(Of("C#")))
	assert.Equal(t, 2, Of("Gb").Diff(Of("Ab")))
	assert.Equal(t, -3, Of("C").Diff(Of("A")))
	assert.Equal(t, 4, Of("D").Diff(Of("F#")))
	assert.Equal(t, -6, Of("F").Diff(Of("B")))
}
