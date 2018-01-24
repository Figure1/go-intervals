go-intervals - The start and end of integer interval management
================================

[![Go Report Card](https://goreportcard.com/badge/github.com/Figure1/go-intervals)](https://goreportcard.com/report/github.com/Figure1/go-intervals) [![GoDoc](https://godoc.org/github.com/Figure1/go-intervals?status.svg)](https://godoc.org/github.com/Figure1/go-intervals)

Go code (golang) package that provides a way to create and manipulate a set of integer intervals.

See it in action:

```go
package main

import (
  "fmt"
  "github.com/Figure1/go-intervals"
)

func main() {
    intervalSet := intervals.New()
    intervalSet.Add(2, 6)
    intervalSet.Add(10, 13)
    intervalSet.add(15, 20)

    intervalSet.Contains(12) // true
    intervalSet.Overlaps(18, 22) // true

    intervalSet.Delete(5, 12)
    fmt.Println(intervalSet) // map[2:4 13:13 18:22]
}
```

------

Installation
============

To install go-intervals, use `go get`:

    go get github.com/Figure1/go-intervals

------

Staying up to date
==================

To update go-intervals to the latest version, use `go get -u github.com/Figure1/go-intervals`.

------

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!