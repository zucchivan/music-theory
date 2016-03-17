# Music 

[![Build Status](https://travis-ci.org/go-music/music.svg?branch=master)](https://travis-ci.org/go-music/music) [![GoDoc](https://godoc.org/github.com/go-music/music?status.svg)](https://godoc.org/github.com/go-music/music)

https://github.com/go-music/music

#### Opinionated models of certain atomic musical concepts

There's an example command-line utility `music.go` to demo the libraries:

    go run music.go -k "C minor 7"
    
This will output:
    
    root: C
    major: false
    minor: true
    tones:
      3: D#
      5: G
      7: A#

