// Package key is an opinionated model of musical signature, by Charney Kaye.
package key

import (
	"testing"

	"github.com/go-music/music/note"
	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	k := Of("C major 7")
	assert.Equal(t, note.C, k.Root)
	assert.Equal(t, note.E, k.Tones[note.I3])
	assert.Equal(t, note.G, k.Tones[note.I5])
	assert.Equal(t, note.B, k.Tones[note.I7])
}

func TestOf_Invalid(t *testing.T) {
	k := Of("P-funk")
	assert.Equal(t, note.NONE, k.Root)
}
