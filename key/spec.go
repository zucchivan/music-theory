// Keys are expressed in readable strings, e.g. C major or Ab minor
package key

import (
	"gopkg.in/yaml.v2"
)

func (k Key) ToYAML() string {
	spec := specFrom(k)
	out, _ := yaml.Marshal(spec)
	return string(out[:])
}

/*
 *
 private */

func specFrom(k Key) specKey {
	s := specKey{}
	s.Root = k.Root.String(k.AdjSymbol)
	s.Mode = k.Mode.String()
	if k.Mode == Major {
		rel := k.RelativeMinor()
		s.Relative.Root = rel.Root.String(k.AdjSymbol)
		s.Relative.Mode = rel.Mode.String()
	} else if k.Mode == Minor {
		rel := k.RelativeMajor()
		s.Relative.Root = rel.Root.String(k.AdjSymbol)
		s.Relative.Mode = rel.Mode.String()
	}
	return s
}

type specKey struct {
	Root string
	Mode string
	Relative specRelativeKey
}

type specRelativeKey struct {
	Root string
	Mode string
}
