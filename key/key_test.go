 // The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
package key

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-music/music/note"
	"github.com/go-music/music/interval"
)

func TestKeys(t *testing.T) {
	for name, expect := range testKeyExpectations {
		k := Of(name)
		assert.Equal(t, expect.Root, k.Root)
		for interval, class := range expect.Tones {
			assert.Equal(t, class, k.Tones[interval])
		}
	}
}

func TestOf_Invalid(t *testing.T) {
	k := Of("P-funk")
	assert.Equal(t, note.NONE, k.Root)
}

type testKeyExpect struct {
	Root note.Class
 	Tones Tones
}

var testKeyExpectations = map[string]testKeyExpect{
	"C": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.E, // major 3rd
			interval.I5: note.G, // perfect 5th
		},
	},

	"C 7": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.E, // major 3rd
			interval.I5: note.G, // perfect 5th
			interval.I7: note.AS, // minor 7th
		},
	},

	"C Seven": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.E, // major 3rd
			interval.I5: note.G, // perfect 5th
			interval.I7: note.AS, // minor 7th
		},
	},

	"C major": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.E, // major 3rd
			interval.I5: note.G, // perfect 5th
		},
	},

	"C maj": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.E, // major 3rd
			interval.I5: note.G, // perfect 5th
		},
	},

	"C major 7": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.E, // major 3rd
			interval.I5: note.G, // perfect 5th
			interval.I7: note.B, // major 7th
		},
	},

	"C minor": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.DS, // minor 3rd
			interval.I5: note.G, // perfect 5th
		},
	},

	"C min": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.DS, // minor 3rd
			interval.I5: note.G, // perfect 5th
		},
	},

	"C minor 7": testKeyExpect{
		Root: note.C,
		Tones: Tones{
			interval.I3: note.DS, // minor 3rd
			interval.I5: note.G, // perfect 5th
			interval.I7: note.AS, // minor 7th
		},
	},

	"XXX": testKeyExpect{
		Root: note.NONE,
	},
}
