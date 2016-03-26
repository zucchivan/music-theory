// The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
package key

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"gopkg.in/music.v0/theory/note"
	"fmt"
)

func TestKeys(t *testing.T) {
	testExpectations := testExpectationManifest{}
	file, err := ioutil.ReadFile("testdata/expectations.yaml")
	assert.Nil(t, err)

	err = yaml.Unmarshal(file, &testExpectations)
	assert.Nil(t, err)

	assert.True(t, len(testExpectations.Keys) > 0)
	for name, expect := range testExpectations.Keys {
		actual := Of(name)
		assert.Equal(t, note.ClassNamed(expect.Root), actual.Root, fmt.Sprintf("name:%v expect.Root:%v actual.Root:%v", name, expect.Root, actual.Root))
		assert.Equal(t, expect.Mode, actual.Mode, fmt.Sprintf("name:%v expect.Mode:%v actual.Mode:%v", name, expect.Mode, actual.Mode))
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
	Root  string
	Mode  KeyMode
}

type testExpectationManifest struct {
	Keys map[string]testKey
}
