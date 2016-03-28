// Note has an adjustment symbol (Sharp or Flat) to render the "accidental notes for a given name (e.g. of a chord, scale or key)
package note

import "regexp"

// AdjSymbolOf the adjustment symbol (Sharp or Flat) for a given name (e.g. of a chord, scale or key)
func AdjSymbolOf(name string) AdjSymbol {
	numSharps := len(rgxSharps.FindAllString(name, -1))
	numFlats := len(rgxFlats.FindAllString(name, -1))
	if numSharps >= numFlats {
		return Sharp
	} else {
		return Flat
	}
}

// Expression of the "accidental notes" as either Sharps or Flats
type AdjSymbol int

const (
	Sharp AdjSymbol = iota
	Flat
)

/*
 *
 private */

var (
	rgxSharps, _ = regexp.Compile("[♯#]")
	rgxFlats, _ = regexp.Compile("[♭b]")
)

