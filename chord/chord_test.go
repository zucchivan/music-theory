// Package chord is an opinionated model of a musical chord, by Charney Kaye.
package chord

// Copyright 2015 Outright Mental, Inc.

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-music/music/key"
	"github.com/go-music/music/note"
)

func TestNewChord(t *testing.T) {
	assert.Equal(t, &Chord{
		Key: key.Of("C minor 7"),
	}, In(key.Of("C minor 7")))

	assert.Equal(t, &Chord{
		Key: key.Of("C major 7"),
	}, In(key.Of("C major 7")))
}

func TestChordNotes(t *testing.T) {
	assert.Equal(t, []*note.Note{
		note.Named("C"),
		note.Named("Ds"),
		note.Named("G"),
		note.Named("As"),
	}, In(key.Of("C minor 7")).Notes())

	assert.Equal(t, []*note.Note{
		note.Named("C"),
		note.Named("E"),
		note.Named("G"),
		note.Named("B"),
	}, In(key.Of("C major 7")).Notes())
}

func TestChordNoteNames(t *testing.T) {
	assert.Equal(t, "C, D#, G, A#", In(key.Of("C minor 7")).NotesCSV())
	assert.Equal(t, "C, E, G, B", In(key.Of("C major 7")).NotesCSV())
}
