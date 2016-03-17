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
	// TODO
}

func TestClassNamed(t *testing.T) {
	// TODO
}

func TestShiftTime(t *testing.T) {
	// TODO
}

func TestFullName(t *testing.T) {
	// TODO
}
