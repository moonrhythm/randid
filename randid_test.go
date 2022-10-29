package randid_test

import (
	"testing"

	"github.com/moonrhythm/randid"
)

func TestGenerate(t *testing.T) {
	id := randid.MustGenerate()
	if len(id) != 32 {
		t.Errorf("len(id) = %d; want 32", len(id))
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randid.MustGenerate()
	}
}
