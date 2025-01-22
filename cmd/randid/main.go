package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/moonrhythm/randid"
)

func main() {
	if len(os.Args) > 1 {
		at := tryParseTime(os.Args[1])
		fmt.Print(randid.At(at))
		return
	}
	fmt.Print(randid.MustGenerate())
}

func tryParseTime(at string) time.Time {
	if t, err := strconv.ParseInt(at, 10, 64); err == nil {
		return time.Unix(t, 0)
	}

	if t, err := time.Parse(time.RFC3339, at); err == nil {
		return t
	}

	if t, err := time.Parse("2006-01-02", at); err == nil {
		return t
	}

	return time.Time{}
}
