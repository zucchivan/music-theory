// A Note is used to represent the relative duration and pitch of a sound.
package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamed(t *testing.T) {
	n := Named("C")
	assert.Equal(t, n, &Note{
		Class: C,
	})
}

func TestOfClass(t *testing.T) {
	n := OfClass(C)
	assert.Equal(t, n, &Note{
		Class: C,
	})
}

func TestClassNamed(t *testing.T) {
	n := ClassNamed("C")
	assert.Equal(t, n, C)
}
