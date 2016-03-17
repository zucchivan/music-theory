// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

// Copyright 2015 Outright Mental, Inc.

import (
	"strings"

	"github.com/go-music/music/key"
	"github.com/go-music/music/note"
	"github.com/go-music/music/interval"
)

// Chord in a particular key
type Chord struct {
	Key key.Key
}

// In the key to create a Chord
func In(k key.Key) *Chord {
	c := &Chord{
		Key: k,
	}
	return c
}

// Notes to obtain the notes from the Chord
func (this *Chord) Notes() (notes []*note.Note) {
	notes = append(notes, note.OfClass(this.Key.Root))
	interval.ForAllIn(this.Key.Tones, func(class note.Class) {
		notes = append(notes, note.OfClass(class))
	})
	return
}

// NotesCSV to return a comma-separated string displaying the chord nicely.
func (this *Chord) NotesCSV() (voice string) {
	notes := this.Notes()
	names := []string{}
	for _, n := range notes {
		names = append(names, string(n.Class))
	}
	return strings.Join(names, ", ")
}
