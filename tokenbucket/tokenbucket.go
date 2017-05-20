package tokenbucket

import (
	"time"
)

// NewTokenBucket returns a new token-bucket.
func NewTokenBucket(rate time.Duration, size uint16) chan struct{} {
	sizeInt := int(size)
	bucket := make(chan struct{}, sizeInt)
	fillBucket(bucket, sizeInt)

	go func(bucket chan<- struct{}) {
		for _ = range time.NewTicker(rate).C {
			fillBucket(bucket, sizeInt)
		}
	}(bucket)

	return bucket
}

func fillBucket(bucket chan<- struct{}, size int) {
	for i := size - len(bucket); i > 0; i-- {
		bucket <- struct{}{}
	}
}
