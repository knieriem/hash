package hash

import (
	"hash"
)

// Hash8 is the common interface implemented by all 8-bit hash functions.
type Hash8 interface {
	hash.Hash
	Sum8() uint8
}

// Hash16 is the common interface implemented by all 16-bit hash functions.
type Hash16 interface {
	hash.Hash
	Sum16() uint16
}
