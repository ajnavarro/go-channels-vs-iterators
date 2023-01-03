# Benchmarks Results

```
goos: linux
goarch: amd64
pkg: github.com/ajnavarro/go-channels-vs-iterators
cpu: AMD Ryzen 5 5600X 6-Core Processor             
BenchmarkChannels-12     	12241628	       106.7 ns/op
BenchmarkIterators-12    	636184224	         1.580 ns/op
PASS
ok  	github.com/ajnavarro/go-channels-vs-iterators	2.605s


```

## Results

- 67,5 times faster using iterators instead of channels.
- No need of extra goroutines.
- Iterators are easier to handle, less things to take into account. Less error prone.
- Con: Iterators are not thread-safe by default, they must be designed to be.
