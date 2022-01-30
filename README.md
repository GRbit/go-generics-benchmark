<h2> go-generics-benchmark </h1>
Different benchmarks for Go generics starting from 1.18

<h3> How to use </h3>

First of all, before Go 1.18 release you have to build
Go 1.18 by yourself or use gotip.
```Bash
$ go install golang.org/dl/gotip@latest
$ gotip download
$ gotip version
go version devel go1.18-a5c0b190 Sun Jan 30 02:29:51 2022 +0000 linux/amd64
```

While **containts_test.go** and **fib_test.go** have simple
go benchmarks which could be run as:
```Bash
# After installing gotip, you can use it instead of 'go' command
$ gotip test -bench=.
```

Benchmarks in the **generated** folder are harder to run.
They are here to test compilation speed with generics and 
with native static-typed code.

To run it, you need to generate new files for every test with 
a new number of functions (**nn** constat in **generate.go**).
In other cases go will save the cache of existing files and compile
time will be decreased.

```Bash
$ sed -i 's/nn *=.*/nn = 10004/' generate.go
$ gotip run generate.go
$ gotip clean --cache 
$ time gotip build generated/generics.go generated/dummy_main.go

real	0m0.520s
user	0m0.864s
sys	0m0.102s
```

You can test different cases by choosing different files
in the following combinations:
```Bash
$ gotip build generated/[generics|static].go generated/[dummy|calls]_main.go
```

The first file choice will determine which version of functions 
will be used: with static types or with generics.
The second choice will determine will all the functions be called
inside of the **func main()**

<h3> Results </h3>

<h4> Simple tests </h4>

BTW: I don't know why the second iteration of the same
tests takes less time. **BenchmarkGenericC1** is equal
to **BenchmarkGenericC2**. If someone can explain
it to me (in a more detailed way than "it's some kind
of optimizations"), please make a PR.

```Bash
$ gotip test -bench=.
goos: linux
goarch: amd64
pkg: tst
cpu: Intel(R) Core(TM) i5-8365U CPU @ 1.60GHz
BenchmarkReflectC1-8     	   20428	     62570 ns/op
BenchmarkGenericC1-8     	 3996031	       300.4 ns/op
BenchmarkNativeC1-8      	 3765502	       300.8 ns/op
BenchmarkReflectC2-8     	   18903	     58921 ns/op
BenchmarkGenericC2-8     	 2190943	       545.4 ns/op
BenchmarkNativeC2-8      	 2074528	       537.8 ns/op
BenchmarkInterfaceF1-8   	   13393	     88923 ns/op
BenchmarkGenericF1-8     	   41120	     30936 ns/op
BenchmarkNativeF1-8      	   37862	     31916 ns/op
BenchmarkInterfaceF2-8   	   13342	     92812 ns/op
BenchmarkGenericF2-8     	   37729	     32275 ns/op
BenchmarkNativeF2-8      	   39980	     32544 ns/op
PASS
ok  	tst	20.541s
```

<h4> Compile time tests </h4>

```Bash
# Static
[0] $ time gotip build static.go dummy_main.go 
real	0m1.347s
user	0m3.566s
sys     0m0.192s
[0] $ time gotip build static.go calls_main.go 
real    0m5.614s
user    0m9.924s
sys     0m0.423s

# Generics
[0] $ time gotip build generics.go dummy_main.go 
real   0m0.499s
user   0m0.843s
sys    0m0.097s
[0] $ time gotip build generics.go calls_main.go 
real   0m5.419s
user   0m10.395s
sys    0m0.409s
```
As you can see, if you don't call generic functions it will
not be compiled at all.
