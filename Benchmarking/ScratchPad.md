To run benchmarking in go, create a file with _test.go suffix and function with Benchmark prefix and argument -> (b *testing.B).  

```
FileName: abc_test.go 
```

```
func BenchmarkAbc(b *testing.B) { 

// code logic 

} 
```

While running via terminal use this command 

```
go test -run none -bench . -benchtime 3s
```

Here, it accepts regex after -run which is none and after -bench which is .(period) means we want to run all benchmarks. Also, we are increasing default benchmark time(1s) to 3s. 

 