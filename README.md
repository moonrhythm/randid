# randid

![Build Status](https://github.com/moonrhythm/randid/actions/workflows/test.yaml/badge.svg?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/moonrhythm/randid)](https://goreportcard.com/report/github.com/moonrhythm/randid)
[![GoDoc](https://pkg.go.dev/badge/github.com/moonrhythm/randid)](https://pkg.go.dev/github.com/moonrhythm/randid)

Random ID generator using timestamp as prefix

## Installation

```bash
go get github.com/moonrhythm/randid
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/moonrhythm/randid"
)

func main() {
    id := randid.MustGenerate().String()
    fmt.Println(id)
}
```

## License

MIT
