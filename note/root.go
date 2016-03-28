// Note can be the Root of a Chord, Key or Scale.
package note

import (
	"regexp"
	"strings"
)

/*
 *
 private */

var (
	rgxSingle, _ = regexp.Compile("^[ABCDEFG]")
	rgxDouble, _ = regexp.Compile("^[ABCDEFG][#♭b]")
)

// Parse all forms using Regexp's against a string
func RootAndRemaining(name string) (root Class, nameWithoutRoot string) {
	if r := rgxDouble.FindString(name); len(r) > 0 {
		root = ClassNamed(r)
		nameWithoutRoot = strings.TrimSpace(name[len(r):])
	} else if r := rgxSingle.FindString(name); len(r) > 0 {
		root = ClassNamed(r)
		nameWithoutRoot = strings.TrimSpace(name[len(r):])
	} else {
		root = Nil
		nameWithoutRoot = name
	}
	return
}
