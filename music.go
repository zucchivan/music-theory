// Opinionated models of the building blocks of music theory
//
// The command-line utility `music.go` will demo the Music libraries:
//
//     go run music.go -k "C minor 7"
//
// This will output:
//
//     root: C
//     major: false
//     minor: true
//     tones:
//       3: D#
//       5: G
//       7: A#
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
