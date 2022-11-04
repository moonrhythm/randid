package randid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moonrhythm/randid"
)

func TestFromString(t *testing.T) {
	id, err := randid.FromString("000102030405060708090a0b0c0d0e0f")
	assert.NoError(t, err)
	assert.Equal(t, id.String(), "000102030405060708090a0b0c0d0e0f")
}

func TestID(t *testing.T) {
	id := randid.ID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	t.Run("String", func(t *testing.T) {
		assert.Equal(t, id.String(), "000102030405060708090a0b0c0d0e0f")
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		x, err := id.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, string(x), `"000102030405060708090a0b0c0d0e0f"`)
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		s := `"000102030405060708090a0b0c0d0e0f"`
		var id randid.ID
		err := id.UnmarshalJSON([]byte(s))
		assert.NoError(t, err)
		assert.Equal(t, id.String(), "000102030405060708090a0b0c0d0e0f")
	})
}
