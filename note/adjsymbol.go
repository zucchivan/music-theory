// Note has an adjustment symbol (Sharp or Flat) to render the "accidental notes for a given name (e.g. of a chord, scale or key)
package note

import "regexp"

// AdjSymbolOf the adjustment symbol (Sharp or Flat) for a given name (e.g. of a chord, scale or key)
func AdjSymbolOf(name string) AdjSymbol {
	numSharps := len(rgxSharpIn.FindAllString(name, -1))
	numFlats := len(rgxFlatIn.FindAllString(name, -1))
	numSharpish := len(rgxSharpishIn.FindAllString(name, -1))
	numFlattish := len(rgxFlattishIn.FindAllString(name, -1))
	// sharp/flat has precedent over sharpish/flattish; overall default is sharp
	if numSharps > 0 && numSharps > numFlats {
		return Sharp
	} else if numFlats > 0 {
		return Flat
	} else if numSharpish > 0 && numSharpish > numFlattish {
		return Sharp
	} else if numFlattish > 0 {
		return Flat
	} else {
		return Sharp
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
	No AdjSymbol = iota
	Sharp
	Flat
)

/*
 *
 private */

var (
	rgxSharpIn, _    = regexp.Compile("[♯#]|major")
	rgxFlatIn, _     = regexp.Compile("[♭b]")
	rgxSharpBegin, _ = regexp.Compile("^[♯#]")
	rgxFlatBegin, _  = regexp.Compile("^[♭b]")
	rgxSharpishIn, _ = regexp.Compile("(M|maj|major|aug)")
	rgxFlattishIn, _ = regexp.Compile("([^a-z]|^)(m|min|minor|dim)")
)
