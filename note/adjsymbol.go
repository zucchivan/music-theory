// Note has an adjustment symbol (Sharp or Flat) to render the "accidental notes for a given name (e.g. of a chord, scale or key)
package note

import "regexp"

// AdjSymbolOf the adjustment symbol (Sharp or Flat) for a given name (e.g. of a chord, scale or key)
func AdjSymbolOf(name string) AdjSymbol {
	numSharps := len(rgxSharpIn.FindAllString(name, -1))
	numFlats := len(rgxFlatIn.FindAllString(name, -1))
	if numSharps >= numFlats {
		return Sharp
	} else {
		return Flat
	}
}

// AdjSymbolBegin the adjustment symbol (Sharp or Flat) that begins a given name (e.g. the Root of a chord, scale or key)
func AdjSymbolBegin(name string) AdjSymbol {
	if rgxSharpBegin.MatchString(name) {
		return Sharp
	} else if rgxFlatBegin.MatchString(name) {
		return Flat
	} else {
		return No
	}
}

// Expression of the "accidental notes" as either Sharps or Flats
type AdjSymbol int

const (
	Sharp AdjSymbol = iota
	Flat
	No
)

/*
 *
 private */

var (
	rgxSharpIn, _ = regexp.Compile("[♯#]")
	rgxFlatIn, _  = regexp.Compile("[♭b]")
	rgxSharpBegin, _ = regexp.Compile("^[♯#]")
	rgxFlatBegin, _  = regexp.Compile("^[♭b]")
)
