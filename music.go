// Opinionated models of the building blocks of music theory
//
// Author:  Charney Kaye
//
// Website:  http://w.charney.io
//
// There's an example command-line utility `music.go` to demo the libraries, with `bin/` wrappers for different functions, as follows.
//
// To build a chord, run `bin/chord` followed by the name of the chord, e.g.:
//
//     $ bin/chord Cm nondominant -5 679
//
//     root: C
//     tones:

//       3: D#
//       6: A
//       7: A#
//       9: D
//
// To list the names of all the known chord-building rules, run `bin/chord ?`:
//
//     $ bin/chord ?
//
//     - Basic
//     - Nondominant
//     - Major Triad
//     - Minor Triad
//     - Augmented Triad
//     - Diminished Triad
//     - Suspended Triad
//     - Omit Fifth
//     - Flat Fifth
//     - Add Sixth
//     - Augmented Sixth
//     - Omit Sixth
//     - Add Seventh
//     - Dominant Seventh
//     - Major Seventh
//     - Minor Seventh
//     - Diminished Seventh
//     - Half Diminished Seventh
//     - Diminished Major Seventh
//     - Augmented Major Seventh
//     - Augmented Minor Seventh
//     - Harmonic Seventh
//     - Omit Seventh
//     - Add Ninth
//     - Dominant Ninth
//     - Major Ninth
//     - Minor Ninth
//     - Sharp Ninth
//     - Omit Ninth
//     - Add Eleventh
//     - Dominant Eleventh
//     - Major Eleventh
//     - Minor Eleventh
//     - Omit Eleventh
//     - Add Thirteenth
//     - Dominant Thirteenth
//     - Major Thirteenth
//     - Minor Thirteenth
//
// Keys
//
// To build a key, use `bin/key` followed by the name of the key, e.g.:
//
//     $ bin/key G#m
//     root: G#
//     mode: Minor
//
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/go-music/music/theory/key"
	"github.com/go-music/music/theory/chord"
)

func init() {
	opt = &Options{}
	opt.Parse()
}

func main() {
	if len(opt.Key) > 0 {
		yaml, _ := yaml.Marshal(key.Of(opt.Key))
		fmt.Printf("%+v", string(yaml[:]))
	}

	if len(opt.Chord) > 0 {
		var out []byte
		if strings.TrimSpace(opt.Chord) == "?" {
			out, _ = yaml.Marshal(chord.ListAllForms())
		} else {
			out, _ = yaml.Marshal(chord.Of(opt.Chord))
		}
		fmt.Printf("%+v", string(out[:]))
	}

	os.Exit(0)
}

var (
	opt *Options
)

type Options struct {
	Key string
	Chord string
}

func (o *Options) Parse() {
	flag.StringVar(&o.Key, "k", "", "Key")
	flag.StringVar(&o.Chord, "c", "", "Chord")
	flag.Parse()
}
