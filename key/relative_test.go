// The relative minor of a major key has the same key signature and starts down a minor third (or equivalently up a major sixth); for example, the relative minor of G major is E minor. Similarly the relative major of a minor key starts up a minor third (or down a major sixth); for example, the relative major of F minor is A♭ major.
package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeMajor(t *testing.T) {
	k := Of("A minor")
	expectRk := Of("C major")
	assert.Equal(t, k.RelativeMajor(), expectRk)
}

func TestRelativeMinor(t *testing.T) {
	k := Of("C major")
	expectRk := Of("A minor")
	assert.Equal(t, k.RelativeMinor(), expectRk)
}
