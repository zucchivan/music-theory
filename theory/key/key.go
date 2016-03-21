// The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
package key

import (
	"github.com/go-music/music/theory/interval"
	"github.com/go-music/music/theory/note"
	"regexp"
)

// Of a particular key, e.g. Of("C minor 7")
func Of(name string) Key {
	k := Key{}
	k.parse(name)
	return k
}

// Key is a model of a musical key signature
type Key struct {
	Root  note.Class
	Mode KeyMode
}

type KeyMode string

const (
	MajorKeyMode = KeyMode("Major")
	MinorKeyMode = KeyMode("Minor")
)

// Tones is a set of pitch classes
type Tones map[interval.Interval]note.Class

/*
 *
 private */

func (this *Key) parse(name string) {
	// parse the root, and keep the remaining string
	this.Root, name = note.RootAndRemaining(name)

	// parse the chord Mode
	switch {
	case rgxKeyMajor.MatchString(name):
		this.Mode = MajorKeyMode
	case rgxKeyMinor.MatchString(name):
		this.Mode = MinorKeyMode
	default:
		this.Mode = MajorKeyMode
	}
}


/*
 *
 private */

var (
	rgxKeyMajor   *regexp.Regexp
	rgxKeyMinor   *regexp.Regexp
)

func init() {
	rgxKeyMajor, _ = regexp.Compile("^(M|maj|major)")
	rgxKeyMinor, _ = regexp.Compile("^(m|min|minor)")
}
