// A perfect octave is the interval between one musical pitch and another with half or double its frequency.
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

//
// Private
//

var rgxOctave *regexp.Regexp

func init() {
	rgxOctave, _ = regexp.Compile("(-*[0-9]+)$")
}
