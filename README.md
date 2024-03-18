# go-coalescefs

Package **coalescefs** provides tools for combining multiple `fs.FS` file-systems into a single coalesced union file-system, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-coalescefs

[![GoDoc](https://godoc.org/github.com/reiver/go-coalescefs?status.svg)](https://godoc.org/github.com/reiver/go-coalescefs)

## Example

Here is an example:
```golang
import "github.com/reiver/go-coalescefs"

// ...

var fs1 fs.FS = os.DirFS("some/path/to/a/directory")
var fs2 fs.FS = os.DirFS("another/path/to/a/directory")

// ...

var filesystem fs.FS = coalescefs.FS(fs1, fs2)

// ...

file, err := filesystem.Open(filename)

```

## Import

To import package **coalescefs** use `import` code like the follownig:
```
import "github.com/reiver/go-coalescefs"
```

## Installation

To install package **coalescefs** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-coalescefs
```

## Author

Package **coalescefs** was written by [Charles Iliya Krempeaux](http://reiver.link)
