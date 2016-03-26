// A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.
package chord

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/go-music/music/theory/note"
	"github.com/go-music/music/theory/interval"
)

func TestFormMatchString(t *testing.T) {

}

func TestChordParseForms(t *testing.T) {
	c := Of("Cm nondominant -5 +6 +7 +9")
	assert.Equal(t, map[interval.Interval]note.Class{
		interval.I3: note.DS,
		interval.I6: note.A,
		interval.I7: note.AS,
		interval.I9: note.D,
	}, c.Tones)
}

/*
 *
 private */

//func assertEquivalentForms(t *testing.T, expectForms []Form, actualForms []Form) {
//	for _, expectForm := range expectForms {
//		assert.Contains(t, actualForms, expectForm,
//			fmt.Sprintf("expect %v in actual %v", expectForms, actualForms))
//	}
//	for _, actualForm := range actualForms {
//		assert.Contains(t, expectForms, actualForm,
//			fmt.Sprintf("actual %v in expected %v", actualForms, expectForms))
//	}
//}
