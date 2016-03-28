// In music, a pitch class (p.c. or pc) is a set of all pitches that are a whole number of octaves apart, e.g., the pitch class C consists of the Cs in all octaves.
package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameOf(t *testing.T) {
	testNameOf(t, "C", C, 0, "C", "C")
	testNameOf(t, "C#", Cs, 0, "C#", "Db")
	testNameOf(t, "Cb", B, -1, "B", "B")
	testNameOf(t, "D", D, 0, "D", "D")
	testNameOf(t, "D#", Ds, 0, "D#", "Eb")
	testNameOf(t, "D♭", Cs, 0, "C#", "Db")
	testNameOf(t, "E", E, 0, "E", "E")
	testNameOf(t, "E#", F, 0, "F", "F")
	testNameOf(t, "E♭", Ds, 0, "D#", "Eb")
	testNameOf(t, "F", F, 0, "F", "F")
	testNameOf(t, "F#", Fs, 0, "F#", "Gb")
	testNameOf(t, "F♭", E, 0, "E", "E")
	testNameOf(t, "G", G, 0, "G", "G")
	testNameOf(t, "G♯", Gs, 0, "G#", "Ab")
	testNameOf(t, "Gb", Fs, 0, "F#", "Gb")
	testNameOf(t, "A", A, 0, "A", "A")
	testNameOf(t, "A#", As, 0, "A#", "Bb")
	testNameOf(t, "Ab", Gs, 0, "G#", "Ab")
	testNameOf(t, "B", B, 0, "B", "B")
	testNameOf(t, "B#", C, 1, "C", "C")
	testNameOf(t, "E♭", Ds, 0, "D#", "Eb")
	testNameOf(t, "Bb", As, 0, "A#", "Bb")
	testNameOf(t, "z", Nil, 0, "-", "-")
	testNameOf(t, "zzzz", Nil, 0, "-", "-")
}

func TestNameStep(t *testing.T) {
	//assert.Equal(t, false, true)
}

func Test_baseNameOf(t *testing.T) {
	assert.Equal(t, C, baseNameOf("C# Major"))
	assert.Equal(t, G, baseNameOf("GbM"))
	assert.Equal(t, Nil, baseNameOf("XXX"))
	assert.Equal(t, Nil, baseNameOf(""))
}

func Test_baseStepOf(t *testing.T) {
	//assert.Equal(t, false, true)
}

func Test_stepFrom(t *testing.T) {
	//assert.Equal(t, false, true)
}

func TestStringOf(t *testing.T) {
	assert.Equal(t, "C#", stringOf(Cs, Sharp))
	assert.Equal(t, "Db", stringOf(Cs, Flat))
	assert.Equal(t, "-", stringOf(Cs, No))
}

/*
 *
 private */

func testNameOf(t *testing.T, name string, expectClass Class, expectOctave Octave, expectStringSharp string, expectStringFlat string) {
	c, o := NameOf(name)
	assert.Equal(t, expectOctave, o)
	assert.Equal(t, expectClass, c)
	assert.Equal(t, expectStringSharp, c.String(Sharp))
	assert.Equal(t, expectStringFlat, c.String(Flat))
}
