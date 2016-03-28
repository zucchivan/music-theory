// Notes, Keys, Chords and Scales
//
// Author:  Charney Kaye
//
// Website:  http://w.charney.io
//
// There's an example command-line utility `music-theory.go` to demo the libraries, with a `bin/` wrapper.
//
// To build and install `music-theory` to your machine:
//
//     make install
//
// Then, to calculate the note pitch classes for a specified Chord:
//
//     $ music-theory chord "Cm nondominant -5 679"
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
//     $ music-theory chords
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
// To determine a key:
//
//     $ music-theory key G#m
//
//     root: G#
//     mode: Minor
//
package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"

	"github.com/go-music-theory/music-theory/chord"
	"github.com/go-music-theory/music-theory/key"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true

	app.Name = "music-theory"
	app.Usage = "Notes, Keys, Chords and Scales"
	app.Version = "0.0.2"
	app.Authors = []cli.Author{cli.Author{Name: "Charney Kaye", Email: "hiya@charney.io"}}
	app.Commands = []cli.Command{

		{ // Build a Chord
			Name:        "chord",
			Aliases:     []string{"c"},
			Usage:       "build a Chord",
			Description: "Chord is a named harmonic set of three or more pitch classes specified by a name, e.g. C or Cm6 or D♭m679-5",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				if len(name) > 0 {
					fmt.Printf("%s", chord.Of(name).ToYAML())
				} else {
					// no arguments
					cli.ShowCommandHelp(c, "chord")
				}
			},
		},

		{ // List all Chords
			Name:        "chords",
			Usage:       "list all known Chords",
			Description: "The Chord DNA is this software is a sequential chain of rules to be executed by matching text in the chord name to its musical implications from the root of the chord.",
			Action: func(c *cli.Context) {
				out, _ := yaml.Marshal(chord.ListAllForms())
				fmt.Printf("%s", string(out[:]))
			},
		},

		{ // Find a Key
			Name:        "key",
			Aliases:     []string{"k"},
			Usage:       "find a Key",
			Description: "The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.",
			Action: func(c *cli.Context) {
				name := c.Args().First()
				if len(name) > 0 {
					out, _ := yaml.Marshal(key.Of(name))
					fmt.Printf("%s", string(out[:]))
				} else {
					// no arguments
					cli.ShowCommandHelp(c, "key")
				}
			},
		},
	}
	app.Run(os.Args)
}
