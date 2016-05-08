// A Note is used to represent the relative duration and pitch of a sound.
package note

import (
	"math"
)

// Diff to another pitch Class calculated in +/- semitones
func (this Class) Diff(targetPitch Class) int {
	diffUp := classDiff(this, targetPitch, 1)
	diffDown := classDiff(this, targetPitch, -1)
	if math.Abs(float64(diffUp)) < math.Abs(float64(diffDown)) {
		return diffUp
	} else {
		return diffDown
	}
}

//
// Private
//

func classDiff(from Class, to Class, inc int) int {
	diff := 0
	for {
		if from == to {
			return diff
		}
		diff += inc
		from, _ = from.Step(inc)
	}
}
