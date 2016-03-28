# Music 

[![Build Status](https://travis-ci.org/go-music/music.svg?branch=master)](https://travis-ci.org/go-music/music) [![GoDoc](https://godoc.org/github.com/go-music/music?status.svg)](https://godoc.org/github.com/go-music/music) [![Go Report Card](https://goreportcard.com/badge/github.com/go-music/music)](https://goreportcard.com/report/github.com/go-music/music) [![codebeat badge](https://codebeat.co/badges/f5c88d18-fff5-4bc6-affb-cdc55488f429)](https://codebeat.co/projects/github-com-go-music-music)

https://github.com/go-music/music

Author: [Charney Kaye](http://w.charney.io)

#### Music theory models in Go

There's an example command-line utility `music-theory.go` to demo the libraries, with a `bin/` wrapper.

To build and install `music-theory` to your machine:

    make install

Then, to calculate the note pitch classes for a specified Chord:

    $ music-theory chord "Cm nondominant -5 679"
    
    root: C
    tones:
      3: D#
      6: A
      7: A#
      9: D

To list the names of all the known chord-building rules, run `bin/chord ?`:

    $ music-theory chords
    
    - Basic
    - Nondominant
    - Major Triad
    - Minor Triad
    - Augmented Triad
    - Diminished Triad
    - Suspended Triad
    - Omit Fifth
    - Flat Fifth
    - Add Sixth
    - Augmented Sixth
    - Omit Sixth
    - Add Seventh
    - Dominant Seventh
    - Major Seventh
    - Minor Seventh
    - Diminished Seventh
    - Half Diminished Seventh
    - Diminished Major Seventh
    - Augmented Major Seventh
    - Augmented Minor Seventh
    - Harmonic Seventh
    - Omit Seventh
    - Add Ninth
    - Dominant Ninth
    - Major Ninth
    - Minor Ninth
    - Sharp Ninth
    - Omit Ninth
    - Add Eleventh
    - Dominant Eleventh
    - Major Eleventh
    - Minor Eleventh
    - Omit Eleventh
    - Add Thirteenth
    - Dominant Thirteenth
    - Major Thirteenth
    - Minor Thirteenth

To determine a key:

    $ music-theory key G#m
    
    root: G#
    mode: Minor

## Note

A Note is used to represent the relative duration and pitch of a sound.

[![GoDoc](https://godoc.org/github.com/go-music/music/theory/note?status.svg)](https://godoc.org/github.com/go-music/music/theory/note) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/theory/note)

## Key

The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.

[![GoDoc](https://godoc.org/github.com/go-music/music/theory/key?status.svg)](https://godoc.org/github.com/go-music/music/theory/key) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/theory/key)

## Chord

A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.

[![GoDoc](https://godoc.org/github.com/go-music/music/theory/chord?status.svg)](https://godoc.org/github.com/go-music/music/theory/chord) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/theory/chord)
