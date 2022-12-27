package gojap

import (
	"github.com/dop251/goja"
)

func RunString(s string) (goja.Value, error) {
	vm := goja.New()
	return vm.RunString("2 + 2")
}

func MustRunString(s string) goja.Value {
	v, err := RunString(s)
	if err != nil {
		panic(err)
	}
	return v
}
