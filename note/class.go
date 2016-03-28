// In music, a pitch class is a set of all pitches that are a whole number of octaves apart, e.g., the pitch class C consists of the Cs in all octaves.
package note

// Class of pitch for a note (across all octaves)
type Class int

const (
	Nil Class = iota
	C
	Cs
	D
	Ds
	E
	F
	Fs
	G
	Gs
	A
	As
	B
)

// NameOf a note will return its Class and Octave
func NameOf(text string) (Class, Octave) {
	return baseNameOf(text).Step(baseStepOf(text))
}

// Step from a class to another class, +/- semitones, +/- octave
func (from Class) Step(inc int) (Class, Octave) {
	return stepFrom(from, inc)
}

// String of the note, expressed with Sharps or Flats
func (from Class) String(with AdjSymbol) string {
	return stringOf(from, with)
}

/*
 *
 private */

func stringOf(from Class, with AdjSymbol) string {
	switch from {
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case A:
		return "A"
	case B:
		return "B"
	}

	if with == Sharp {
		return stringSharpOf(from)
	} else if with == Flat {
		return stringFlatOf(from)
	}

	return "-"
}

func stringSharpOf(from Class) string {
	switch from {
	case Cs:
		return "C#"
	case Ds:
		return "D#"
	case Fs:
		return "F#"
	case Gs:
		return "G#"
	case As:
		return "A#"
	}
	return "-"
}

func stringFlatOf(from Class) string {
	switch from {
	case Cs:
		return "Db"
	case Ds:
		return "Eb"
	case Fs:
		return "Gb"
	case Gs:
		return "Ab"
	case As:
		return "Bb"
	}
	return "-"
}

func baseNameOf(text string) Class {
	if len(text) > 0 {
		switch text[0:1] {
		case "C":
			return C
		case "D":
			return D
		case "E":
			return E
		case "F":
			return F
		case "G":
			return G
		case "A":
			return A
		case "B":
			return B
		default:
			return Nil
		}
	} else {
		return Nil
	}
}

func baseStepOf(text string) int {
	if len(text) < 2 {
		return 0
	}

	switch AdjSymbolBegin(text[1:]) {
	case Sharp:
		return 1
	case Flat:
		return -1
	default:
		return 0
	}
}

func stepFrom(name Class, inc int) (Class, Octave) {
	if inc > 0 {
		return stepFromUp(name, inc)
	} else if inc < 0 {
		return stepFromDown(name, -inc)
	}
	return name, 0
}

func stepFromUp(name Class, inc int) (Class, Octave) {
	octave := Octave(0)
	for i := 0; i < inc; i++ {
		shift := stepUp[name]
		name = shift.Name
		octave += shift.Octave
	}
	return name, octave
}

func stepFromDown(name Class, inc int) (Class, Octave) {
	octave := Octave(0)
	for i := 0; i < inc; i++ {
		shift := stepDown[name]
		name = shift.Name
		octave += shift.Octave
	}
	return name, octave
}

type step struct {
	Name   Class
	Octave Octave
}

var (
	stepUp = map[Class]step{
		Nil: step{Nil, 0},
		C:   step{Cs, 0},
		Cs:  step{D, 0},
		D:   step{Ds, 0},
		Ds:  step{E, 0},
		E:   step{F, 0},
		F:   step{Fs, 0},
		Fs:  step{G, 0},
		G:   step{Gs, 0},
		Gs:  step{A, 0},
		A:   step{As, 0},
		As:  step{B, 0},
		B:   step{C, 1},
	}
	stepDown = map[Class]step{
		Nil: step{Nil, 0},
		C:   step{B, -1},
		Cs:  step{C, 0},
		D:   step{Cs, 0},
		Ds:  step{D, 0},
		E:   step{Ds, 0},
		F:   step{E, 0},
		Fs:  step{F, 0},
		G:   step{Fs, 0},
		Gs:  step{G, 0},
		A:   step{Gs, 0},
		As:  step{A, 0},
		B:   step{As, 0},
	}
)
