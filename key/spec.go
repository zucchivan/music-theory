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
	return s
}

type specKey struct {
	Root  string
	Mode string
}
