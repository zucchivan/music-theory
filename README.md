# Music 

[![Build Status](https://travis-ci.org/go-music/music.svg?branch=master)](https://travis-ci.org/go-music/music) [![GoDoc](https://godoc.org/github.com/go-music/music?status.svg)](https://godoc.org/github.com/go-music/music) [![Go Report Card](https://goreportcard.com/badge/github.com/go-music/music)](https://goreportcard.com/report/github.com/go-music/music)

https://github.com/go-music/music

#### Music theory models in Go

There's an example command-line utility `music.go` to demo the libraries, with some helpful wrappers to run it from the command line.

To build a chord, use `bin/chord`:

    $ bin/chord Cm nondominant -5 679
    
This will output:
    
    root: C
    tones:
      3: D#
      6: A
      7: A#
      9: D
      
Use use `bin/key` to build a key:

    $ bin/key G#m
    root: G#
    mode: Minor

Author: [Charney Kaye](http://w.charney.io)

## Note

A Note is used to represent the relative duration and pitch of a sound.

[![GoDoc](https://godoc.org/github.com/go-music/music/theory/note?status.svg)](https://godoc.org/github.com/go-music/music/theory/note) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/theory/note)

## Interval

An interval is the difference between two pitches.

[![GoDoc](https://godoc.org/github.com/go-music/music/theory/interval?status.svg)](https://godoc.org/github.com/go-music/music/theory/interval) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/theory/interval)

## Key

The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.

[![GoDoc](https://godoc.org/github.com/go-music/music/theory/key?status.svg)](https://godoc.org/github.com/go-music/music/theory/key) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/theory/key)

## Chord

A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.

[![GoDoc](https://godoc.org/github.com/go-music/music/theory/chord?status.svg)](https://godoc.org/github.com/go-music/music/theory/chord) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/theory/chord)
