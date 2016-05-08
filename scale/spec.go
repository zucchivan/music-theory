// Scales are expressed in readable strings, e.g. CMb5b7 or Cm679-5
package scale

import (
	"gopkg.in/yaml.v2"
)

func (c Scale) ToYAML() string {
	spec := specFrom(c)
	out, _ := yaml.Marshal(spec)
	return string(out[:])
}

//
// Private
//

func specFrom(c Scale) specScale {
	s := specScale{}
	s.Root = c.Root.String(c.AdjSymbol)
	s.Tones = make(map[int]string)
	for i, t := range c.Tones {
		s.Tones[int(i)] = t.String(c.AdjSymbol)
	}
	return s
}

type specScale struct {
	Root  string
	Tones map[int]string
}
