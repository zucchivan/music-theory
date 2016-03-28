// The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
package key

import (
	"github.com/go-music-theory/music-theory/note"
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
	Root note.Class
	Mode Mode
}

// Mode is the mode of a key, e.g. Major or Minor
type Mode string

const (
	MajorKeyMode Mode = "Major"
	MinorKeyMode Mode = "Minor"
)

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
	rgxKeyMajor *regexp.Regexp
	rgxKeyMinor *regexp.Regexp
)

func init() {
	rgxKeyMajor, _ = regexp.Compile("^(M|maj|major)")
	rgxKeyMinor, _ = regexp.Compile("^(m|min|minor)")
}
