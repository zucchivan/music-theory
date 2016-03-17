// The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.
package key

import (
	"strings"

	"github.com/go-music/music/note"
	"github.com/go-music/music/interval"
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
	Major bool
	Minor bool
	//
	Tones Tones
}

// Tones is a set of pitch classes
type Tones map[interval.Interval]note.Class

/*
 *
 private */

func (this *Key) parse(name string) {
	this.Tones = make(map[interval.Interval]note.Class)

	// split key name into chunks
	chunks := strings.Split(name, " ")

	// shift the first chunk, which is required to be the root class
	chunkRoot, chunks := chunks[0], chunks[1:]
	this.Root = note.ClassNamed(chunkRoot)
	if this.Root == note.NONE {
		return
	}

	// parse the rest of the chunks
	for _, chunk := range chunks {
		this.parseChunk(chunk)
	}

	// parse any implications, given all known information
	this.parseImplications()
}

func (this *Key) parseChunk(chunk string) {
	chunk = strings.ToLower(chunk)
	switch {
	case rgxKeyMajor.MatchString(chunk):
		this.Major = true
		this.Tones[interval.I3], _ = this.Root.Step(4) // major 3rd
	case rgxKeyMinor.MatchString(chunk):
		this.Minor = true
		this.Tones[interval.I3], _ = this.Root.Step(3) // minor 3rd
	case rgxKeySeventh.MatchString(chunk):
		if this.Major {
			this.Tones[interval.I7], _ = this.Root.Step(11) // major 7th
		} else if this.Minor {
			this.Tones[interval.I7], _ = this.Root.Step(10) // minor 7th
		} else {
			this.Tones[interval.I7], _ = this.Root.Step(10) // minor 7th
		}
	}
}

func (this *Key) parseImplications() {
	// perfect 5th, if none exists
	if _, exists := this.Tones[interval.I5]; !exists {
		this.Tones[interval.I5], _ = this.Root.Step(7)
	}

	// major 3rd, if none exists
	if _, exists := this.Tones[interval.I3]; !exists {
		this.Tones[interval.I3], _ = this.Root.Step(4)
	}
}

/*
 *
 private */

var (
	rgxKeyMajor *regexp.Regexp
	rgxKeyMinor *regexp.Regexp
	rgxKeySeventh *regexp.Regexp
)

func init() {
	rgxKeyMajor, _ = regexp.Compile("^maj")
	rgxKeyMinor, _ = regexp.Compile("^min")
	rgxKeySeventh, _ = regexp.Compile("^7|sev")
}
