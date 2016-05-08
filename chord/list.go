// It's possible to export a list of all known chord parsing rules.
package chord

import (
	"gopkg.in/yaml.v2"
)

type List []string

// ToYAML any List to an array of strings
func (l List) ToYAML() string {
	out, _ := yaml.Marshal(l)
	return string(out[:])
}

var ChordFormList List

//
// Private
//

func init() {
	for _, f := range forms {
		ChordFormList = append(ChordFormList, f.Name)
	}
}
