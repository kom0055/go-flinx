# go-flinx

## Functional LINQ IN Go

[![Sourcegraph](https://sourcegraph.com/github.com/kom0055/go-flinx/-/badge.svg)](https://sourcegraph.com/github.com/kom0055/go-flinx?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/kom0055/go-flinx)
[![Build Status](https://travis-ci.org/kom0055/go-flinx.svg?branch=master)](https://travis-ci.org/kom0055/go-flinx)
[![codecov](https://codecov.io/gh/kom0055/go-flinx/branch/master/graph/badge.svg)](https://codecov.io/gh/kom0055/go-flinx)
[![rcard](https://goreportcard.com/badge/github.com/kom0055/go-flinx)](https://goreportcard.com/report/github.com/kom0055/go-flinx)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/kom0055/go-flinx/master/LICENSE)

Thanks [ahmetb/go-linq](https://github.com/ahmetb/go-linq)  for a guidance.

Rewrite [ahmetb/go-linq](https://github.com/ahmetb/go-linq)  for these reasons:

1. Not compatible with generics. I've been enough to use type assertions.
2. Low performance when use generic function. You could use generic through function ends with 't', but it's not that
   efficient and not that elegant.

So to solve my problems, I choose to rewrite it using generics and functional-programing.
You know that generics in Go is not that flexible, but feature of Currying and Lazy Evaluation could make this.

## Installation

When used with Go modules, use the following import path:

    go get github.com/kom0055/go-flinx

## Quickstart

If you were a dotNet/C# developer but a gopher now, you must miss the smooth grammar of LINQ.

If you were a fans of functional-programing but a gopher now, you must be eager to code in functional style.

Now, you could both enjoy LINQ and functional-programing through this package.

A simple case like below

```go
squares := ToSlice(Select(func (x int) int { return x * x })(Range(1, 10)))
```

For more cases, you could visit [example_test.go](example_test.go)

## Performance

Using ahmetb's go-linq comparison benchmark (see [benchmark_test.go](benchmark_test.go), test size was 10000000):

SelectWhereFirst:

| Lib                      |          Time           |
|:-------------------------|:-----------------------:|
| go-flinx                 |       153.4 ns/op       |
| go-linq                  |       162.6 ns/op       |
| gp-linq(generics func)   |       1043 ns/op        |

Sum:

| Lib                      |       Time       |
|:-------------------------|:----------------:|
| go-flinx                 |  31092617 ns/op  |
| go-linq                  | 119863792 ns/op  |
| gp-linq(generics func)   | 2249057541 ns/op |

ZipSkipTake:

| Lib                      |    Time     |
|:-------------------------|:-----------:|
| go-flinx                 | 120.3 ns/op |
| go-linq                  | 108.4 ns/op |
| gp-linq(generics func)   | 562.8 ns/op |

FromChannel:

| Lib                      |       Time       |
|:-------------------------|:----------------:|
| go-flinx                 | 1206086167 ns/op |
| go-linq                  | 1312612709 ns/op |
| gp-linq(from channelt)   | 1975103125 ns/op |

We could make conclusion that if you chase for efficient performance and elegant grammar, go-flinx is a good choice. 
