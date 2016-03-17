// The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
package key

import (
	"testing"

	"github.com/go-music/music/note"
	"github.com/stretchr/testify/assert"
)

func TestOf_C(t *testing.T) {
	k := Of("C")
	assert.Equal(t, note.C, k.Root)
	assert.Equal(t, note.E, k.Tones[note.I3]) // major 3rd
	assert.Equal(t, note.G, k.Tones[note.I5]) // perfect 5th
}

func TestOf_C_7(t *testing.T) {
	k := Of("C 7")
	assert.Equal(t, note.C, k.Root)
	assert.Equal(t, note.E, k.Tones[note.I3]) // major 3rd
	assert.Equal(t, note.G, k.Tones[note.I5]) // perfect 5th
	assert.Equal(t, note.AS, k.Tones[note.I7]) // minor 7th
}

func TestOf_C_major(t *testing.T) {
	k := Of("C major")
	assert.Equal(t, note.C, k.Root)
	assert.Equal(t, note.E, k.Tones[note.I3]) // major 3rd
	assert.Equal(t, note.G, k.Tones[note.I5]) // perfect 5th
}

func TestOf_C_major_7(t *testing.T) {
	k := Of("C major 7")
	assert.Equal(t, note.C, k.Root)
	assert.Equal(t, note.E, k.Tones[note.I3]) // major 3rd
	assert.Equal(t, note.G, k.Tones[note.I5]) // perfect 5th
	assert.Equal(t, note.B, k.Tones[note.I7]) // major 7th
}

func TestOf_C_minor(t *testing.T) {
	k := Of("C minor")
	assert.Equal(t, note.C, k.Root)
	assert.Equal(t, note.DS, k.Tones[note.I3]) // minor 3rd
	assert.Equal(t, note.G, k.Tones[note.I5]) // perfect 5th
}

func TestOf_C_minor_7(t *testing.T) {
	k := Of("C minor 7")
	assert.Equal(t, note.C, k.Root)
	assert.Equal(t, note.DS, k.Tones[note.I3]) // minor 3rd
	assert.Equal(t, note.G, k.Tones[note.I5]) // perfect 5th
	assert.Equal(t, note.AS, k.Tones[note.I7]) // minor 7th
}

func TestOf_Invalid(t *testing.T) {
	k := Of("P-funk")
	assert.Equal(t, note.NONE, k.Root)
}
