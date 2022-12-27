[![Tests on Linux, MacOS and Windows](https://github.com/bep/gojap/workflows/Test/badge.svg)](https://github.com/bep/gojap/actions?query=workflow:Test)
[![Go Report Card](https://goreportcard.com/badge/github.com/bep/gojap)](https://goreportcard.com/report/github.com/bep/gojap)
[![GoDoc](https://godoc.org/github.com/bep/gojap?status.svg)](https://godoc.org/github.com/bep/gojap)

- [Baseline benchmarks with new VM per execution and no compile cache](#baseline-benchmarks-with-new-vm-per-execution-and-no-compile-cache)
- [Benchmarks using a compilation cache and VM pool](#benchmarks-using-a-compilation-cache-and-vm-pool)

This is a small library that wraps [goja](https://github.com/dop251/goja/) and adds a VM pool and compilation cache. The primary intended use case for this is fast execution of a small set of small JavaScript programs with lots of different data contexts.

## Baseline benchmarks with new VM per execution and no compile cache

```
Benchmark branch "main"

 go version go1.19.3 darwin/arm64

goos: darwin
goarch: arm64
pkg: github.com/bep/gojap
BenchmarkRunString/no_args/Baseline/serial-10  	    8080	    146926 ns/op	  292696 B/op	    3629 allocs/op
BenchmarkRunString/no_args/Baseline/parallel-10         	    7945	    147632 ns/op	  292705 B/op	    3629 allocs/op
BenchmarkRunString/one_arg/Baseline/serial-10           	    7335	    152738 ns/op	  292696 B/op	    3629 allocs/op
BenchmarkRunString/one_arg/Baseline/parallel-10         	    7830	    146396 ns/op	  292697 B/op	    3629 allocs/op
BenchmarkRunString/big_script/Baseline/serial-10        	    1453	    807831 ns/op	  392055 B/op	    5219 allocs/op
BenchmarkRunString/big_script/Baseline/parallel-10      	    4455	    266412 ns/op	  392035 B/op	    5219 allocs/op
PASS
ok  	github.com/bep/gojap	8.238s


name                                       time/op
RunString/no_args/Baseline/serial-10       147µs ± 0%
RunString/no_args/Baseline/parallel-10     148µs ± 0%
RunString/one_arg/Baseline/serial-10       153µs ± 0%
RunString/one_arg/Baseline/parallel-10     146µs ± 0%
RunString/big_script/Baseline/serial-10    808µs ± 0%
RunString/big_script/Baseline/parallel-10  266µs ± 0%

name                                       alloc/op
RunString/no_args/Baseline/serial-10       293kB ± 0%
RunString/no_args/Baseline/parallel-10     293kB ± 0%
RunString/one_arg/Baseline/serial-10       293kB ± 0%
RunString/one_arg/Baseline/parallel-10     293kB ± 0%
RunString/big_script/Baseline/serial-10    392kB ± 0%
RunString/big_script/Baseline/parallel-10  392kB ± 0%

name                                       allocs/op
RunString/no_args/Baseline/serial-10       3.63k ± 0%
RunString/no_args/Baseline/parallel-10     3.63k ± 0%
RunString/one_arg/Baseline/serial-10       3.63k ± 0%
RunString/one_arg/Baseline/parallel-10     3.63k ± 0%
RunString/big_script/Baseline/serial-10    5.22k ± 0%
RunString/big_script/Baseline/parallel-10  5.22k ± 0%
```

## Benchmarks using a compilation cache and VM pool

```
 go version go1.19.3 darwin/arm64

goos: darwin
goarch: arm64
pkg: github.com/bep/gojap
BenchmarkRunString/no_args/Cached/serial-10         	13894812	        74.72 ns/op	      32 B/op	       1 allocs/op
BenchmarkRunString/no_args/Cached/parallel-10       	 8473996	       141.2 ns/op	      38 B/op	       1 allocs/op
BenchmarkRunString/one_arg/Cached/serial-10         	 5172565	       226.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkRunString/one_arg/Cached/parallel-10       	 7516302	       161.8 ns/op	      37 B/op	       1 allocs/op
BenchmarkRunString/big_script/Cached/serial-10      	   99742	     11929 ns/op	    2323 B/op	     280 allocs/op
BenchmarkRunString/big_script/Cached/parallel-10    	  420262	      2557 ns/op	    2332 B/op	     280 allocs/op
PASS
ok  	github.com/bep/gojap	7.840s


name                                     time/op
RunString/no_args/Cached/serial-10       74.7ns ± 0%
RunString/no_args/Cached/parallel-10      141ns ± 0%
RunString/one_arg/Cached/serial-10        227ns ± 0%
RunString/one_arg/Cached/parallel-10      162ns ± 0%
RunString/big_script/Cached/serial-10    11.9µs ± 0%
RunString/big_script/Cached/parallel-10  2.56µs ± 0%

name                                     alloc/op
RunString/no_args/Cached/serial-10        32.0B ± 0%
RunString/no_args/Cached/parallel-10      38.0B ± 0%
RunString/one_arg/Cached/serial-10        32.0B ± 0%
RunString/one_arg/Cached/parallel-10      37.0B ± 0%
RunString/big_script/Cached/serial-10    2.32kB ± 0%
RunString/big_script/Cached/parallel-10  2.33kB ± 0%
```
