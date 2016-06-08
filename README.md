# Anyhash

[![GoDoc](https://godoc.org/github.com/dlsniper/anyhash?status.svg)](https://godoc.org/github.com/dlsniper/anyhash) [![Go Report Card](https://goreportcard.com/badge/github.com/dlsniper/anyhash)](https://goreportcard.com/report/github.com/dlsniper/anyhash) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

Anyhash is an attempt to hash anything you throw at it and get a slice of bytes.

It relies on the wonderful [xxHash](https://github.com/OneOfOne/xxhash) algorithm to do it fast.

Please read the whole readme before you proceed.

## Usage

```go
func demo() {
    hash := anyhash.Hash("my data goes here")
    // your code here
}
```

## Benchmarks

Because that's what's important:

- CPU
```
Native

BenchmarkSimpleJSON-8            1000000              1371 ns/op
BenchmarkSimpleDump-8            1000000              1846 ns/op
BenchmarkFullJSON-8               200000              8032 ns/op
BenchmarkFullDump-8               500000              3237 ns/op
BenchmarkTagsJSON-8              2000000               780 ns/op
BenchmarkTagsDump-8              1000000              1249 ns/op


cgo

BenchmarkSimpleJSON-8            1000000              1353 ns/op
BenchmarkSimpleDump-8            1000000              2074 ns/op
BenchmarkFullJSON-8               200000              8094 ns/op
BenchmarkFullDump-8               500000              3429 ns/op
BenchmarkTagsJSON-8              2000000               784 ns/op
BenchmarkTagsDump-8              1000000              1626 ns/op
```

- Memory
```
Native

BenchmarkSimpleJSON-8            1000000              1400 ns/op             328 B/op          3 allocs/op
BenchmarkSimpleDump-8            1000000              2034 ns/op             576 B/op         16 allocs/op
BenchmarkFullJSON-8               200000              9961 ns/op            2632 B/op         12 allocs/op
BenchmarkFullDump-8               500000              4042 ns/op            1104 B/op         20 allocs/op
BenchmarkTagsJSON-8              2000000               926 ns/op             184 B/op          2 allocs/op
BenchmarkTagsDump-8              1000000              1550 ns/op             296 B/op         12 allocs/op


cgo

BenchmarkSimpleJSON-8            1000000              1360 ns/op             328 B/op          3 allocs/op
BenchmarkSimpleDump-8            1000000              2165 ns/op             624 B/op         18 allocs/op
BenchmarkFullJSON-8               200000              8398 ns/op            2632 B/op         12 allocs/op
BenchmarkFullDump-8               500000              3432 ns/op            1152 B/op         22 allocs/op
BenchmarkTagsJSON-8              2000000               760 ns/op             184 B/op          2 allocs/op
BenchmarkTagsDump-8              1000000              1521 ns/op             344 B/op         14 allocs/op
```

## If you've made it this far...

Please don't use it as dependency, it's awful, just copy-paste the code if you must.

I've done it in < 5 minutes, so can you!

### Production usage

Don't do it, just don't.

## Thanks

Thanks to the folks on slack for improving this, you know who you are.

## Credits

Credits go to [cnf](https://github.com/cnf) for the test data which happens to be compatible with [this package](https://github.com/cnf/structhash)
