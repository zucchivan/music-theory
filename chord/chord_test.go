// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/go-music-theory/music-theory/key"
	"github.com/go-music-theory/music-theory/note"
)

func TestChordExpectations(t *testing.T) {
	testExpectations := testExpectationManifest{}
	file, err := ioutil.ReadFile("testdata/expectations.yaml")
	assert.Nil(t, err)

	err = yaml.Unmarshal(file, &testExpectations)
	assert.Nil(t, err)

	assert.True(t, len(testExpectations.Chords) > 0)
	for name, expect := range testExpectations.Chords {
		actual := Of(name)
		assert.Equal(t, note.ClassNamed(expect.Root), actual.Root, fmt.Sprintf("name:%v expect.Root:%v actual.Root:%v", name, note.ClassNamed(expect.Root), actual.Root))
		for i, c := range expect.Tones {
			assert.Equal(t, note.ClassNamed(c), actual.Tones[i], fmt.Sprintf("name:%v expect.Tones[%v]:%v actual.Tones[%v]:%v", name, i, c, i, actual.Tones[i]))
		}
		for i, c := range actual.Tones {
			assert.Equal(t, note.ClassNamed(expect.Tones[i]), c, fmt.Sprintf("name:%v actual.Tones[%v]:%v expect.Tones[%v]:%v", name, i, c, i, expect.Tones[i]))
		}
	}
}

func TestNotes(t *testing.T) {
	c := Of("Cm nondominant -5 +6 +7 +9")
	assert.Equal(t, []*note.Note{
		&note.Note{Class: note.Ds},
		&note.Note{Class: note.A},
		&note.Note{Class: note.As},
		&note.Note{Class: note.D},
	}, c.Notes())
}

func TestOf_Invalid(t *testing.T) {
	k := key.Of("P-funk")
	assert.Equal(t, note.Nil, k.Root)
}

/*
 *
 private */

type testKey struct {
	Root  string
	Tones map[Interval]string
}

type testExpectationManifest struct {
	Chords map[string]testKey
}
