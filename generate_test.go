package randid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moonrhythm/randid"
)

func TestGenerate(t *testing.T) {
	assert.NotPanics(t, func() {
		randid.MustGenerate()
	})
	assert.NotEmpty(t, randid.MustGenerate())
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randid.MustGenerate()
	}
}
