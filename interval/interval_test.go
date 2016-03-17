// An interval is the difference between two pitches.
package interval

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-music/music/note"
)

func TestInterval(t *testing.T) {
	assert.Equal(t, 15, len(Order))
}

func TestForAllIn(t *testing.T) {
	tones := map[Interval]note.Class{
		I2: note.DS,
		I5: note.G,
		I7: note.AS,
	}
	ForAllIn(tones, func(class note.Class) {
		assert.NotEmpty(t, class)
	})
}
