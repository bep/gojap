package gojap

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRunString(t *testing.T) {
	c := qt.New(t)
	e := New()
	c.Assert(e.MustRunString("2 + 2").ToInteger(), qt.Equals, int64(4))
}

func BenchmarkRunString(b *testing.B) {
	b.Run("no args", func(b *testing.B) {
		e := New()
		for i := 0; i < b.N; i++ {
			e.MustRunString("2 + 2")
		}
	})
}
