package gojap

import (
	"sync"

	"github.com/dop251/goja"
)

var vmPool = sync.Pool{
	New: func() interface{} {
		return goja.New()
	},
}

func getVm() *goja.Runtime {
	return vmPool.Get().(*goja.Runtime)
}

func putVm(vm *goja.Runtime) {
	vmPool.Put(vm)
}

// New returns a new Exec.
func New() *Exec {
	return &Exec{
		pcache: make(map[string]*goja.Program),
	}
}

// Exec is a JavaScript executor that caches compiled programs.
type Exec struct {
	pcache   map[string]*goja.Program
	pcacheMu sync.RWMutex
}

// RunStringc compiles and runs the given string s as a JavaScript program.
// Note that the compiled program is cached, so any script needs to
// is compiled using strict mode to prevent global pollution.
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

	vm := getVm()
	defer putVm(vm)
	return vm.RunProgram(p)
}

func (e *Exec) MustRunString(s string) goja.Value {
	v, err := e.RunString(s)
	if err != nil {
		panic(err)
	}
	return v
}
