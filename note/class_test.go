// In music, a pitch class (p.c. or pc) is a set of all pitches that are a whole number of octaves apart, e.g., the pitch class C consists of the Cs in all octaves.
package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameOf(t *testing.T) {
	testNameOf(t, "C", C, 0)
	testNameOf(t, "C#", Cs, 0)
	testNameOf(t, "Cb", B, -1)
	testNameOf(t, "D", D, 0)
	testNameOf(t, "Dsharp", Ds, 0)
	testNameOf(t, "Dflat", Cs, 0)
	testNameOf(t, "E", E, 0)
	testNameOf(t, "Esh", F, 0)
	testNameOf(t, "Efl", Ds, 0)
	testNameOf(t, "F", F, 0)
	testNameOf(t, "Fs", Fs, 0)
	testNameOf(t, "Ff", E, 0)
	testNameOf(t, "G", G, 0)
	testNameOf(t, "G#", Gs, 0)
	testNameOf(t, "Gb", Fs, 0)
	testNameOf(t, "A", A, 0)
	testNameOf(t, "A#", As, 0)
	testNameOf(t, "Ab", Gs, 0)
	testNameOf(t, "B", B, 0)
	testNameOf(t, "B#", C, 1)
	testNameOf(t, "E♭", Ds, 0)
	testNameOf(t, "Bb", As, 0)
	testNameOf(t, "z", Nil, 0)
	testNameOf(t, "zzzz", Nil, 0)
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

/*
 *
 private */

func testNameOf(t *testing.T, name string, expectClass Class, expectOctave Octave) {
	n, o := NameOf(name)
	assert.Equal(t, expectOctave, o)
	assert.Equal(t, expectClass, n)
}
