// Note can be the Root of a Chord, Key or Scale.
package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootOf(t *testing.T) {
	assertRootAndRemaining(t, "C", C, "")
	assertRootAndRemaining(t, "Cmaj", C, "maj")
	assertRootAndRemaining(t, "B♭min", As, "min")
	assertRootAndRemaining(t, "C#dim", Cs, "dim")
	assertRootAndRemaining(t, "JAMS", Nil, "JAMS")
}

/*
 *
 private */

func assertRootAndRemaining(t *testing.T, fromChord string, expectRoot Class, expectRemaining string) {
	root, remaining := RootAndRemaining(fromChord)
	assert.Equal(t, expectRoot, root)
	assert.Equal(t, expectRemaining, remaining)
}
