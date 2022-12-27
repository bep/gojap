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

type Arg struct {
	Name  string
	Value any
}

// RunString compiles and runs the given string s as a JavaScript program.
// Note that the compiled program is cached using the string s as the key.
func (e *Exec) RunString(s string, args ...Arg) (goja.Value, error) {
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
	defer func() {
		for _, arg := range args {
			vm.GlobalObject().Delete(arg.Name)
		}
		putVm(vm)
	}()

	for _, arg := range args {
		if err := vm.Set(arg.Name, arg.Value); err != nil {
			return nil, err
		}
	}

	return vm.RunProgram(p)
}

func (e *Exec) MustRunString(s string, args ...Arg) goja.Value {
	v, err := e.RunString(s, args...)
	if err != nil {
		panic(err)
	}
	return v
}
