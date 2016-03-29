// Key has a Mode, e.g. Major or Minor
package key

import (
	"regexp"
)

// Mode is the mode of a key, e.g. Major or Minor
type Mode int

const (
	Nil Mode = iota
	Major
	Minor
)

// String of the Mode, e.g. "Major" or "Minor"
func (of Mode) String() string {
	switch of {
	case Nil:
		return "Nil"
	case Major:
		return "Major"
	case Minor:
		return "Minor"
	}
	return ""
}

/*
 *
 private */

var (
	rgxMajor, _ = regexp.Compile("^(M|maj|major)")
	rgxMinor, _ = regexp.Compile("^(m\\b|min|minor|Minor)")
)

func (k *Key) parseMode(name string) {
	// parse the chord Mode
	k.Mode = modeOf(name)
}

func modeOf(name string) Mode {
	switch {
	case rgxMinor.MatchString(name):
		return Minor
	case rgxMajor.MatchString(name):
		return Major
	default:
		return Major
	}
}
