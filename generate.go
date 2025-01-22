package randid

import (
	"crypto/rand"
	"encoding/binary"
	"io"
	"sync"
	"time"
)

const poolSize = 16 * 16

var (
	pool    [poolSize]byte
	poolPos = poolSize
	poolMu  sync.Mutex
)

func Generate() (ID, error) {
	var b [16]byte

	binary.BigEndian.PutUint64(b[:], uint64(time.Now().UnixNano()))

	poolMu.Lock()
	if poolPos == poolSize {
		_, err := io.ReadFull(rand.Reader, pool[:])
		if err != nil {
			poolMu.Unlock()
			return ID{}, err
		}
		poolPos = 0
	}
	copy(b[8:], pool[poolPos:poolPos+8])
	poolPos += 8
	poolMu.Unlock()

	return b, nil
}

func MustGenerate() ID {
	id, err := Generate()
	if err != nil {
		panic(err)
	}
	return id
}

func GenerateAt(t time.Time) (ID, error) {
	var b [16]byte = At(t)

	poolMu.Lock()
	if poolPos == poolSize {
		_, err := io.ReadFull(rand.Reader, pool[:])
		if err != nil {
			poolMu.Unlock()
			return ID{}, err
		}
		poolPos = 0
	}
	copy(b[8:], pool[poolPos:poolPos+8])
	poolPos += 8
	poolMu.Unlock()

	return b, nil
}

func At(t time.Time) ID {
	var b [16]byte
	binary.BigEndian.PutUint64(b[:], uint64(t.UnixNano()))
	return b
}
