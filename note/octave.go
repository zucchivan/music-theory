// Package note is an opinionated model of a musical note, by Charney Kaye.
package note

import (
	"regexp"
	"strconv"
)

// Octave models a musical octave
type Octave int

// Octave of text returns a new Octave
func OctaveOf(text string) Octave {
	o, _ := strconv.Atoi(rgxOctave.FindString(text))
	return Octave(o)
}

/*
 *
 private */

var rgxOctave *regexp.Regexp

func init() {
	rgxOctave, _ = regexp.Compile("(-*[0-9]+)$")
}
