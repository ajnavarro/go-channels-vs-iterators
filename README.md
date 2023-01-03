# Benchmarks Results

## Channels:

```
goos: linux
goarch: amd64
pkg: github.com/ajnavarro/go-channels-vs-iterators
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkChannels
BenchmarkChannels-12
 2634987	       497.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ajnavarro/go-channels-vs-iterators	1.795s

```

## Iterators:

```
goos: linux
goarch: amd64
pkg: github.com/ajnavarro/go-channels-vs-iterators
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkIterators
BenchmarkIterators-12    	38726305	        29.41 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/ajnavarro/go-channels-vs-iterators	1.178s

```

## Results

- 17 times faster using iterators instead of channels.
- No need of extra goroutines.
- Iterators are easier to handle, less things to take into account. Less error prone.
- Con: Iterators are not thread-safe by default, they must be designed to be.