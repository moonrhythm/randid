package main

import (
	"fmt"

	"github.com/moonrhythm/randid"
)

func main() {
	fmt.Print(randid.MustGenerate())
}
