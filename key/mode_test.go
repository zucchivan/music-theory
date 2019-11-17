// Key has a Mode, e.g. Major or Minor
package key

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"testing"
)

func TestModeString(t *testing.T) {
	assert.Equal(t, "Major", Major.String())
	assert.Equal(t, "Minor", Minor.String())
	assert.Equal(t, "Nil", Nil.String())
	m := Mode(5)
	assert.Equal(t, "", m.String())
}

func TestModeOf(t *testing.T) {
	assert.Equal(t, Major, modeOf("Major"))
	assert.Equal(t, Major, modeOf("M"))
	assert.Equal(t, Major, modeOf("major"))

	assert.Equal(t, Minor, modeOf("minor"))
	assert.Equal(t, Minor, modeOf("min"))
	assert.Equal(t, Minor, modeOf("m"))

	assert.Equal(t, Major, modeOf("joe"))
}
