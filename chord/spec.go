// Chords are expressed in readable strings, e.g. CMb5b7 or Cm679-5
package chord

import (
	"gopkg.in/yaml.v2"
)

func (c Chord) ToYAML() string {
	spec := specFrom(c)
	out, _ := yaml.Marshal(spec)
	return string(out[:])
}

//
// Private
//

func specFrom(c Chord) specChord {
	s := specChord{}
	s.Root = c.Root.String(c.AdjSymbol)
	s.Tones = make(map[int]string)
	for i, t := range c.Tones {
		s.Tones[int(i)] = t.String(c.AdjSymbol)
	}
	return s
}

type specChord struct {
	Root  string
	Tones map[int]string
}
