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
	for { //put token
		if len(bucket) < sizeInt {
			bucket <- token
			continue
		}
		break
	}

	go func(bucket chan<- byte) {
		tick := time.NewTicker(rate)
		for _ = range tick.C {
			if len(bucket) < sizeInt {
				bucket <- token
			}
		}
	}(bucket)

	return bucket
}
