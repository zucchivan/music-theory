// It's possible to export a list of all known scale parsing rules.
package scale

import (
	"gopkg.in/yaml.v2"
)

type List []string

// ToYAML any List to an array of strings
func (l List) ToYAML() string {
	out, _ := yaml.Marshal(l)
	return string(out[:])
}

var ScaleModeList List

//
// Private
//

func init() {
	for _, f := range modes {
		ScaleModeList = append(ScaleModeList, f.Name)
	}
}
