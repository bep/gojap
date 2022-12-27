package gojap

import (
	"sync"

	"github.com/dop251/goja"
)

func New() *Exec {
	return &Exec{
		pcache: make(map[string]*goja.Program),
	}
}

type Exec struct {
	pcache   map[string]*goja.Program
	pcacheMu sync.RWMutex
}

// RunStringc compiles and runs the given string s as a JavaScript program.
// Note that the compiled program is cached, so any script needs to
// use strict mode to prevent global pollution.
func (e *Exec) RunString(s string) (goja.Value, error) {
	e.pcacheMu.RLock()
	p, ok := e.pcache[s]
	e.pcacheMu.RUnlock()
	if !ok {
		var err error
		p, err = goja.Compile("", s, true)
		if err != nil {
			return nil, err
		}
		e.pcacheMu.Lock()
		e.pcache[s] = p
		e.pcacheMu.Unlock()
	}

	vm := goja.New()
	return vm.RunProgram(p)
}

func (e *Exec) MustRunString(s string) goja.Value {
	v, err := e.RunString(s)
	if err != nil {
		panic(err)
	}
	return v
}
