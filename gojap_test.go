package gojap

import (
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/dop251/goja"
	qt "github.com/frankban/quicktest"
)

func TestRunString(t *testing.T) {
	c := qt.New(t)
	e := New()
	c.Assert(e.MustRunString("2 + 2").ToInteger(), qt.Equals, int64(4))
	c.Assert(e.MustRunString("2 + 2 + k", Arg{"k", 32}).ToInteger(), qt.Equals, int64(36))
}

func TestRunStringParallel(t *testing.T) {
	c := qt.New(t)
	e := New()

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			c.Assert(e.MustRunString("2 + 2 + k", Arg{"k", i}).ToInteger(), qt.Equals, int64(4+i))
			if i%2 == 0 {
				c.Assert(e.MustRunString("2 + 2 + l", Arg{"l", i}).ToInteger(), qt.Equals, int64(4+i))
			} else {
				// Make sure l hasn't leaked across the cached vms.
				_, err := e.RunString("2+l")
				c.Assert(err, qt.IsNotNil)
			}
		}()
	}

	wg.Wait()

}

func BenchmarkRunString(b *testing.B) {

	baseline := func(s string, args ...Arg) goja.Value {
		p, err := goja.Compile("", s, true)
		if err != nil {
			b.Fatal(err)
		}
		vm := goja.New()

		for _, arg := range args {
			if err := vm.Set(arg.Name, arg.Value); err != nil {
				b.Fatal(err)
			}
		}

		v, err := vm.RunProgram(p)
		if err != nil {
			b.Fatal(err)
		}
		return v

	}

	runbfn := func(b *testing.B, fn func(s string, args ...Arg) goja.Value, s string, args ...Arg) {
		b.Run("serial", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(s, args...)
			}
		})

		b.Run("parallel", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					fn(s, args...)
				}
			})
		})

	}

	runb := func(b *testing.B, s string, args ...Arg) {
		b.Run("Baseline", func(b *testing.B) {
			runbfn(b, baseline, s, args...)
		})
		b.Run("Cached", func(b *testing.B) {
			e := New()
			runbfn(b, e.MustRunString, s, args...)
		})
	}

	b.Run("no args", func(b *testing.B) {
		runb(b, "2 + 2")
	})

	b.Run("one arg", func(b *testing.B) {
		runb(b, "2 + 2", Arg{"k", 32})
	})

	b.Run("many scripts", func(b *testing.B) {
		e := New()

		for i := 0; i < b.N; i++ {
			j := i % 100
			s := "2 + 2 +" + strconv.Itoa(j)
			e.MustRunString(s)
		}
	})

	b.Run("big script", func(b *testing.B) {
		k := Arg{"k", 32}
		s := strings.Repeat("2 + 2 + k + ", 100)
		s = s[:len(s)-3]
		b.ResetTimer()
		runb(b, s, k)

	})

}
