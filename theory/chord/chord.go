// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

import (
	"gopkg.in/music.v0/theory/interval"
	"gopkg.in/music.v0/theory/note"
)

// Chord in a particular key
type Chord struct {
	Root       note.Class
	Tones      map[interval.Interval]note.Class
}

// Of a particular key, e.g. Of("C minor 7")
func Of(name string) Chord {
	c := Chord{}
	c.parse(name)
	return c
}

// Notes to obtain the notes from the Chord
func (this *Chord) Notes() (notes []*note.Note) {
	interval.ForAllIn(this.Tones, func(class note.Class) {
		notes = append(notes, note.OfClass(class))
	})
	return
}

/*
 *
 private */

func (this *Chord) parse(name string) {
	this.Tones = make(map[interval.Interval]note.Class)

	// parse the root, and keep the remaining string
	this.Root, name = note.RootAndRemaining(name)

	// parse the chord Form
	this.parseForms(name)
}
