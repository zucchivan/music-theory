// Note has an adjustment symbol (Sharp or Flat) to render the "accidental notes for a given name (e.g. of a chord, scale or key)
package note

import (
	"testing"

	"gopkg.in/stretchr/testify.v1/assert"
)

func TestAdjSymbolOf(t *testing.T) {
	assert.Equal(t, Sharp, AdjSymbolOf("C"))
	assert.Equal(t, Flat, AdjSymbolOf("CMb5b7"))
	assert.Equal(t, Sharp, AdjSymbolOf("C#"))
	assert.Equal(t, Flat, AdjSymbolOf("Gb"))
	assert.Equal(t, Flat, AdjSymbolOf("G♭M"))
	assert.Equal(t, Sharp, AdjSymbolOf("A#m"))
	assert.Equal(t, Sharp, AdjSymbolOf("A♯M♯5"))
	assert.Equal(t, Flat, AdjSymbolOf("C minor"))
	assert.Equal(t, Flat, AdjSymbolOf("C dim"))
	assert.Equal(t, Sharp, AdjSymbolOf("CM M9 m7")) // More Sharpish than Flattish
	assert.Equal(t, Flat, AdjSymbolOf("Cm m9 M7"))  // More Flattish than Sharpish
	assert.Equal(t, Sharp, AdjSymbolOf("C major"))
}

func TestAdjSymbolBegin(t *testing.T) {
	assert.Equal(t, No, AdjSymbolBegin("C"[1:]))
	assert.Equal(t, No, AdjSymbolBegin("CMb5b7"[1:]))
	assert.Equal(t, Sharp, AdjSymbolBegin("C#"[1:]))
	assert.Equal(t, Flat, AdjSymbolBegin("Gb"[1:]))
	assert.Equal(t, Flat, AdjSymbolBegin("G♭M"[1:]))
	assert.Equal(t, Sharp, AdjSymbolBegin("A#m"[1:]))
	assert.Equal(t, Sharp, AdjSymbolBegin("A♯M♯5"[1:]))
}
