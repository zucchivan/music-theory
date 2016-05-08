// Note can be the Root of a Chord, Key or Scale.
package note

import (
	"regexp"
	"strings"
)

//
// Private
//

var (
	rgxSingle, _ = regexp.Compile("^[ABCDEFG]")
	rgxDouble, _ = regexp.Compile("^[ABCDEFG][♯#♭b]")
)

// Parse all forms using Regexp's against a string
func RootAndRemaining(name string) (Class, string) {
	if r := rgxDouble.FindString(name); len(r) > 0 {
		return ClassNamed(r), strings.TrimSpace(name[len(r):])
	}

	if r := rgxSingle.FindString(name); len(r) > 0 {
		return ClassNamed(r), strings.TrimSpace(name[len(r):])
	}

	return Nil, name
}
