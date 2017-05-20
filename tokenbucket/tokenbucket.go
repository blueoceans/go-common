package tokenbucket

import (
	"time"
)

const (
	token byte = 00
)

// NewTokenBucket returns a new token-bucket.
func NewTokenBucket(rate time.Duration, size uint16) chan byte {
	sizeInt := int(size)
	bucket := make(chan byte, sizeInt)
	fillBucket(bucket, sizeInt)

	go func(bucket chan<- byte) {
		tick := time.NewTicker(rate)
		for _ = range tick.C {
			fillBucket(bucket, sizeInt)
		}
	}(bucket)

	return bucket
}

func fillBucket(bucket chan<- byte, size int) {
	for i := size - len(bucket); i > 0; i-- {
		bucket <- token
	}
}
