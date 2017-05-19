package tokenbucket

import (
	"time"
)

var (
	token = Token{}
)

// Token is a bit in the token-bucket.
type Token struct {
}

// NewTokenBucket returns a new token-bucket.
func NewTokenBucket(rate time.Duration, size uint16) chan Token {
	sizeInt := int(size)
	bucket := make(chan Token, sizeInt)
	for { //put token
		if len(bucket) < sizeInt {
			bucket <- token
			continue
		}
		break
	}

	go func(bucket chan<- Token) {
		tick := time.NewTicker(rate)
		for _ = range tick.C {
			if len(bucket) < sizeInt {
				bucket <- token
			}
		}
	}(bucket)

	return bucket
}
