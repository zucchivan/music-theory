// Opinionated models of the building blocks of music theory
//
// There's an example command-line utility `music.go` to demo the libraries, with some helpful wrappers to run it from the command line.
//
// To build a chord, use `bin/chord`:
//
//     $ bin/chord Cm nondominant -5 679
//
// This will output:
//
//     root: C
//     tones:
//     3: D#
//     6: A
//     7: A#
//     9: D
//
// Use use `bin/key` to build a key:
//
//     $ bin/key G#m
//     root: G#
//     mode: Minor
//
// Author:  Charney Kaye
// Website:  http://w.charney.io
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
