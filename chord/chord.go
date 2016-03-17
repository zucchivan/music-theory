// Package chord is an opinionated model of a musical chord, by Charney Kaye.
package chord

// Copyright 2015 Outright Mental, Inc.

import (
	"strings"

	"github.com/go-music/music/key"
	"github.com/go-music/music/note"
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
	for _, interval := range note.IntervalOrder {
		if val, ok := this.Key.Tones[interval]; ok {
			notes = append(notes, note.OfClass(val))
		}
	}
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
