// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/go-music/music/theory/interval"
	"github.com/go-music/music/theory/key"
	"github.com/go-music/music/theory/note"
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
		assert.Equal(t, expect.Root, actual.Root, fmt.Sprintf("name:%v expect.Root:%v actual.Root:%v", name, expect.Root, actual.Root))
		for interval, class := range expect.Tones {
			assert.Equal(t, note.ClassNamed(class), actual.Tones[interval], fmt.Sprintf("name:%v expect.Tones[%v]:%v actual.Tones[%v]:%v", name, interval, class, interval, actual.Tones[interval]))
		}
		for interval, class := range actual.Tones {
			assert.Equal(t, note.ClassNamed(expect.Tones[interval]), class, fmt.Sprintf("name:%v actual.Tones[%v]:%v expect.Tones[%v]:%v", name, interval, class, interval, expect.Tones[interval]))
		}
	}
}

func TestNotes(t *testing.T) {
	c := Of("Cm nondominant -5 +6 +7 +9")
	assert.Equal(t, []*note.Note{
		&note.Note{Class:note.DS},
		&note.Note{Class:note.A},
		&note.Note{Class:note.AS},
		&note.Note{Class:note.D},
	}, c.Notes())
}

func TestOf_Invalid(t *testing.T) {
	k := key.Of("P-funk")
	assert.Equal(t, note.NONE, k.Root)
}

/*
 *
 private */

type testKey struct {
	Root  note.Class
	Tones map[interval.Interval]string
}

type testExpectationManifest struct {
	Chords map[string]testKey
}
