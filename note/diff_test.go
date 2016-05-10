// A Note is used to represent the relative duration and pitch of a sound.
package note

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClass_Diff(t *testing.T) {
	assert.Equal(t, 5, Cs.Diff(Fs))
	assert.Equal(t, -5, Fs.Diff(Cs))
	assert.Equal(t, 2, Gs.Diff(As))
	assert.Equal(t, -3, C.Diff(A))
	assert.Equal(t, 4, D.Diff(Fs))
	assert.Equal(t, -6, F.Diff(B))
}

func TestClass_Diff_PanicFromNil(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()
	Nil.Diff(Fs)
}
