// Key Difference to another Key can be calculated in +/- semitones
package key

// Diff to another Key calculated in +/- semitones
func (this Key) Diff(targetKey Key) int {
	return this.Root.Diff(targetKey.Root)
}
