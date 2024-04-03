# go-iolsep

Package **iolsep** provides a limited-reader that reads a single line from an io.Reader, for the Go programming language.

A line is terminated by one of these:

* `"\n"`
* `"\u0085"`
* `"\u2028"`
* `"\u2029"`
* `io.EOF`

Note that `"\r"` (which would include `"\r\n"` and `"\n\r"`) is intentionally missing from this list.
But it still works out.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-iolsep

[![GoDoc](https://godoc.org/github.com/reiver/go-iolsep?status.svg)](https://godoc.org/github.com/reiver/go-iolsep)

## Usage

```golang
var readcloser io.ReadCloser = iolsep.NewLineReadCloser(reader)
```

## Import

To import package **iolsep** use `import` code like the follownig:
```
import "github.com/reiver/go-iolsep"
```

## Installation

To install package **iolsep** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-iolsep
```

## Author

Package **iolsep** was written by [Charles Iliya Krempeaux](http://reiver.link)
