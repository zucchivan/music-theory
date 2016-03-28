// In music, a pitch class is a set of all pitches that are a whole number of octaves apart, e.g., the pitch class C consists of the Cs in all octaves.
package note

// Class of pitch for a note (across all octaves)
type Class string

const (
	NONE Class = "-"
	C    Class = "C"
	CS   Class = "C#"
	D    Class = "D"
	DS   Class = "D#"
	E    Class = "E"
	F    Class = "F"
	FS   Class = "F#"
	G    Class = "G"
	GS   Class = "G#"
	A    Class = "A"
	AS   Class = "A#"
	B    Class = "B"
)

// NameOf a note will return its Class and Octave
func NameOf(text string) (Class, Octave) {
	return baseNameOf(text).Step(baseStepOf(text))
}

// Step from a class to another class, +/- semitones, +/- octave
func (from Class) Step(inc int) (Class, Octave) {
	return stepFrom(from, inc)
}

/*
 *
 private */

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
			return NONE
		}
	} else {
		return NONE
	}
}

func baseStepOf(text string) (inc int) {
	if len(text) < 2 {
		return 0
	} else if text[1:2] == "#" { // Sharp e.g. "C#"
		return 1
	} else if text[1:2] == "s" {
		// Sharp e.g. "Cs" or "Csharp"
		return 1
	} else if text[1:2] == "b" { // Flat e.g. "Db"
		return -1
	} else if text[1:2] == "f" { // Flat e.g. "Cf" or "Cflat"
		return -1
	} else if len(text) >= 4 && text[1:4] == "♭" { // Flat e.g. "E♭" (special character length=3)
		return -1
	} else {
		return 0
	}
}

func stepFrom(name Class, inc int) (Class, Octave) {
	octave := Octave(0)
	if inc > 0 {
		for i := 0; i < inc; i++ {
			shift := stepUp[name]
			name = shift.Name
			octave += shift.Octave
		}
	} else if inc < 0 {
		for i := 0; i > inc; i-- {
			shift := stepDown[name]
			name = shift.Name
			octave += shift.Octave
		}
	}
	return name, octave
}

type step struct {
	Name   Class
	Octave Octave
}

var (
	stepUp = map[Class]step{
		NONE: step{NONE, 0},
		C:    step{CS, 0},
		CS:   step{D, 0},
		D:    step{DS, 0},
		DS:   step{E, 0},
		E:    step{F, 0},
		F:    step{FS, 0},
		FS:   step{G, 0},
		G:    step{GS, 0},
		GS:   step{A, 0},
		A:    step{AS, 0},
		AS:   step{B, 0},
		B:    step{C, 1},
	}
	stepDown = map[Class]step{
		NONE: step{NONE, 0},
		C:    step{B, -1},
		CS:   step{C, 0},
		D:    step{CS, 0},
		DS:   step{D, 0},
		E:    step{DS, 0},
		F:    step{E, 0},
		FS:   step{F, 0},
		G:    step{FS, 0},
		GS:   step{G, 0},
		A:    step{GS, 0},
		AS:   step{A, 0},
		B:    step{AS, 0},
	}
)
