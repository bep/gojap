package gojap

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRunString(t *testing.T) {
	c := qt.New(t)
	c.Assert(MustRunString("2 + 2").ToInteger(), qt.Equals, int64(4))
}

func BenchmarkRunString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustRunString("2 + 2")
	}
}
