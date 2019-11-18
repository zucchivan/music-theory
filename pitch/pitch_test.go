package pitch

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

// table with proper values can be found here https://en.wikipedia.org/wiki/Scientific_pitch_notation

func TestOfClassAndOctave(t *testing.T) {
	assertPitchOfClassAndOctave(t, "440.00Hz", "A", "4")
	assertPitchOfClassAndOctave(t, "32.70Hz", "C", "1")
	assertPitchOfClassAndOctave(t, "92.50Hz", "Gb", "2")
	assertPitchOfClassAndOctave(t, "1244.51Hz", "D♯", "6")
	assertPitchOfClassAndOctave(t, "31608.53Hz", "B", "10")
	assertPitchOfClassAndOctave(t, "8.18Hz", "C", "-1")
}

func TestOfNote(t *testing.T) {
	assertPitchOfNote(t, "440.00Hz", "A4")
	assertPitchOfNote(t, "32.70Hz", "C1")
	assertPitchOfNote(t, "92.50Hz", "Gb2")
	assertPitchOfNote(t, "1244.51Hz", "D#6")
	assertPitchOfNote(t, "31608.53Hz", "B10")
	assertPitchOfNote(t, "8.18Hz", "C-1")
}

func assertPitchOfClassAndOctave(t *testing.T, expected string, class string, octave string) {
	actual, err := OfClassAndOctave(class, octave)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func assertPitchOfNote(t *testing.T, expected string, name string) {
	actual, err := OfNote(name)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
