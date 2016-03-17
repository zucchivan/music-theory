 // The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
package key

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/go-music/music/note"
)

func TestKeys(t *testing.T) {
	testExpectations := testExpectationManifest{}
	file, err := ioutil.ReadFile("testdata/expectations.yaml")
	assert.Nil(t, err)

	err = yaml.Unmarshal(file, &testExpectations)
	assert.Nil(t, err)

	assert.True(t, len(testExpectations.Keys) > 0)
	for name, expect := range testExpectations.Keys {
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

/*
 *
 private */

type testKey struct {
	Root  note.Class
	Tones Tones
}

type testExpectationManifest struct {
	Keys map[string]testKey
}
