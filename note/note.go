// A Note is used to represent the relative duration and pitch of a sound.
package note

// Note models a musical note
type Note struct {
	Performer string  // name of performer (for reference)
	Position  float64 // >0, # of beats
	Duration  float64 // >0, # of beats
	Class     Class   // Class of pitch
	Octave    Octave  // Octave #
	Code      string  // custom, directly from music
}

// Named note returns a Note model
func Named(text string) (n *Note) {
	n = &Note{}

	// First the name, including octave shift.
	n.Class, n.Octave = NameOf(text)

	// Last, add the originally named octave.
	n.Octave += OctaveOf(text)

	return
}

// OfClass pitch returns a Note model
func OfClass(class Class) (n *Note) {
	n = &Note{}
	n.Class = class
	return
}

// ClassNamed returns a pitch Class
func ClassNamed(text string) Class {
	n := &Note{}

	// First the name, including octave shift.
	n.Class, n.Octave = NameOf(text)

	// Last, add the originally named octave.
	n.Octave += OctaveOf(text)

	return n.Class
}

// ShiftTime to copy the note to another position in time.
func (from *Note) ShiftTime(t float64) *Note {
	to := *from
	from.Position += t
	return &to
}
