package btree

import (
	"testing"
)

func TestBucket_EncodeDecode(t *testing.T) {
	buffer := make([]byte, bucketSize)

	LoadBucket(0, buffer)
}
