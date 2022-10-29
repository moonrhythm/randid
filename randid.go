package randid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
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

func Generate() (string, error) {
	var b [16]byte

	binary.BigEndian.PutUint64(b[:], uint64(time.Now().UnixNano()))

	poolMu.Lock()
	if poolPos == poolSize {
		_, err := io.ReadFull(rand.Reader, pool[:])
		if err != nil {
			poolMu.Unlock()
			return "", err
		}
		poolPos = 0
	}
	copy(b[8:], pool[poolPos:poolPos+8])
	poolPos += 8
	poolMu.Unlock()

	var s [32]byte
	hex.Encode(s[:], b[:])
	return string(s[:]), nil
}

func MustGenerate() string {
	id, err := Generate()
	if err != nil {
		panic(err)
	}
	return id
}
