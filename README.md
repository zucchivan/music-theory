[![Build Status](https://travis-ci.org/go-music-theory/music-theory.svg?branch=master)](https://travis-ci.org/go-music-theory/music-theory) [![GoDoc](https://godoc.org/gopkg.in/music-theory.v0?status.svg)](https://godoc.org/gopkg.in/music-theory.v0) [![Go Report Card](https://goreportcard.com/badge/gopkg.in/music-theory.v0)](https://goreportcard.com/report/gopkg.in/music-theory.v0) [![codebeat badge](https://codebeat.co/badges/2636c257-5ea9-47dd-8194-871e29178c46)](https://codebeat.co/projects/github-com-go-music-theory-music-theory)

# Music Theory

https://gopkg.in/music-theory.v0

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

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/note?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/note) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/gopkg.in/music-theory.v0/note)

## Key

The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/key?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/key) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/gopkg.in/music-theory.v0/key)

## Chord

A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.

[![GoDoc](https://godoc.org/gopkg.in/music-theory.v0/chord?status.svg)](https://godoc.org/gopkg.in/music-theory.v0/chord) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/gopkg.in/music-theory.v0/chord)
