// Scales have different Modes, such as Triad, Seventh, Extended, Added/Omitted, Specific or General.
package scale

import (
	"regexp"
)

// Mode is identified by positive/negative regular expressions, and then adds/removes pitch classes by interval from the root of the scale.
type Mode struct {
	Name string
	pos  *regexp.Regexp
	set  ModeIntervals
	omit ModeOmit
}

// ModeAdd maps an interval-from-scale-root to a +/1 semitone adjustment
type ModeIntervals []int

// ModeOmit maps an interval-from-scale-root to omit
type ModeOmit []Interval

// MatchString processes the positive/negative regular expressions to determine if this mode matches a string.
func (this *Mode) MatchString(s string) bool {
	return this.matchPosNegString(s)
}

//
// Private
//

// Regular expression to use mid-word, gluing together mode expression parts
var nExp = "[. ]*"

// Regular expressions for different utilities
var (
	majorExp       = "(M|maj|major)"
	minorStrictExp = "([^a-z ]|^)(m|min|minor)"
	minorExp       = "(m|min|minor)"

	//flatExp  = "(f|flat|b|♭)"
	//sharpExp = "(#|s|sharp)"
	//halfExp  = "half"

	//omitExp = "(omit|\\-)"

	naturalExp    = "(nat|natural)"
	melodicExp    = "(mel|melodic)"
	ascendExp     = "(asc|ascend)"
	descendExp    = "(desc|descend)"
	diminishedExp = "(dim|dimin|diminished)"
	augmentedExp  = "(aug|augment|augmented)"
	harmonicExp   = "(harm|harmonic)"
	//dominantExp    = "(^|dom|dominant)"
	//nondominantExp = "(non|nondom|nondominant)"
	//suspendedExp   = "(sus|susp|suspend|suspended)"

	locrianExp    = "(loc|locrian)"
	ionianExp     = "(ion|ionian)"
	dorianExp     = "(dor|dorian)"
	phrygianExp   = "(phr|phrygian)"
	lydianExp     = "(lyd|lydian)"
	mixolydianExp = "(mix|mixolydian)"
	aeolianExp    = "(aeo|aeolian)"

	ionianIntervals     = ModeIntervals{2, 2, 1, 2, 2, 2}
	dorianIntervals     = ModeIntervals{2, 1, 2, 2, 2, 1}
	phrygianIntervals   = ModeIntervals{1, 2, 2, 2, 1, 2}
	lydianIntervals     = ModeIntervals{2, 2, 2, 1, 2, 2}
	mixolydianIntervals = ModeIntervals{2, 2, 1, 2, 2, 1}
	aeolianIntervals    = ModeIntervals{2, 1, 2, 2, 1, 2}
	locrianIntervals    = ModeIntervals{1, 2, 2, 1, 2, 2}
)

// modes is an ordered set of rules to match, and corresponding scale intervals to setup.
var modes = []Mode{

	// Basic

	Mode{
		Name: "Default (Major)",
		set:  ionianIntervals,
	},

	Mode{
		Name: "Minor",
		pos:  exp(minorStrictExp),
		set:  aeolianIntervals,
	},

	Mode{
		Name: "Major",
		pos:  exp(majorExp),
		set:  ionianIntervals,
	},

	Mode{
		Name: "Natural Minor",
		pos:  exp(naturalExp + nExp + minorExp),
		set:  aeolianIntervals,
	},

	Mode{
		Name: "Diminished",
		pos:  exp(diminishedExp),
		set:  ModeIntervals{2, 1, 2, 1, 2, 1, 2},
	},

	Mode{
		Name: "Augmented",
		pos:  exp(augmentedExp),
		set:  ModeIntervals{3, 1, 3, 1, 3},
		omit: ModeOmit{I7},
	},

	Mode{
		Name: "Melodic Minor Ascend",
		pos:  exp(melodicExp + nExp + minorExp + nExp + ascendExp),
		set:  ModeIntervals{2, 1, 2, 2, 2, 2},
	},

	Mode{
		Name: "Melodic Minor Descend",
		pos:  exp(melodicExp + nExp + minorExp + nExp + descendExp),
		set:  ModeIntervals{2, 1, 2, 2, 1, 2},
	},

	Mode{
		Name: "Harmonic Minor",
		pos:  exp(harmonicExp + nExp + minorExp),
		set:  ModeIntervals{2, 1, 2, 2, 1, 3},
	},

	Mode{
		Name: "Ionian",
		pos:  exp(ionianExp),
		set:  ionianIntervals,
	},

	Mode{
		Name: "Dorian",
		pos:  exp(dorianExp),
		set:  dorianIntervals,
	},

	Mode{
		Name: "Phrygian",
		pos:  exp(phrygianExp),
		set:  phrygianIntervals,
	},

	Mode{
		Name: "Lydian",
		pos:  exp(lydianExp),
		set:  lydianIntervals,
	},

	Mode{
		Name: "Mixolydian",
		pos:  exp(mixolydianExp),
		set:  mixolydianIntervals,
	},

	Mode{
		Name: "Aeolian",
		pos:  exp(aeolianExp),
		set:  aeolianIntervals,
	},

	Mode{
		Name: "Locrian",
		pos:  exp(locrianExp),
		set:  locrianIntervals,
	},
}

func exp(s string) *regexp.Regexp {
	r, _ := regexp.Compile(s)
	return r
}

func (this *Mode) matchPosNegString(s string) bool {
	if this.pos == nil {
		//fmt.Printf("[%s] matched %s by default", s, this.Name)
		return true
	} else if this.pos.MatchString(s) {
		//log.Printf("[%s] matched pos %s\n", s, this.Name)
		//if this.neg != nil {
		//	if this.neg.MatchString(s) {
		//		return false
		//	}
		//}
		return true
	} else {
		return false
	}
}

// Build the scale by processing all Modes against the given name.
func (this *Scale) parseModes(name string) {
	var toDelete []Interval
	for _, f := range modes {
		if f.MatchString(name) {
			toDelete = append(toDelete, this.applyMode(f)...)
		}
	}
	for _, t := range toDelete {
		delete(this.Tones, t)
	}
	return
}

func (this *Scale) applyMode(f Mode) (toDelete []Interval) {
	ct := I1
	this.Tones[ct] = this.Root
	for _, c := range f.set {
		ct++
		this.Tones[ct], _ = this.Tones[ct-1].Step(c)
	}
	for _, t := range f.omit {
		toDelete = append(toDelete, t)
	}
	return
}
