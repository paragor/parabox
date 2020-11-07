package btree

import (
	"bytes"
	"encoding/gob"
)

const (
	noParent     = -1
	reservedData = 2 ^ 10
	bucketSize   = 2 ^ 14
)

type Meta struct {
	isLeaf bool
	parent int64
}

type Bucket struct {
	start        int64
	payloadStart int64
	data         []byte
	cachedMeta   *Meta
}

func LoadBucket(start int64, data []byte) *Bucket {
	return &Bucket{start: start, payloadStart: start + reservedData, data: data}
}

func (b Bucket) IsLeaf() (bool, error) {
	meta, err := b.fetchMeta()
	if err != nil {
		return false, err
	}

	return meta.isLeaf, nil
}

func (b Bucket) getParentId() (int64, error) {
	meta, err := b.fetchMeta()
	if err != nil {
		return 0, err
	}

	return meta.parent, nil
}

func (b Bucket) fetchMeta() (*Meta, error) {
	if b.cachedMeta == nil {
		meta, err := b.fetchMetaForce()
		if err != nil {
			return nil, err
		}
		b.cachedMeta = meta
	}

	return b.cachedMeta, nil
}

func (b Bucket) fetchMetaForce() (*Meta, error) {
	buffer := make([]byte, reservedData)
	decoder := gob.NewDecoder(bytes.NewBuffer(buffer))
	var meta Meta
	err := decoder.Decode(&meta)
	if err != nil {
		return nil, err
	}

	return &meta, nil
}
func (b Bucket) saveMeta(meta Meta) error {
	buffer := make([]byte, reservedData)

	encoder := gob.NewEncoder(bytes.NewBuffer(buffer))
	err := encoder.Encode(meta)
	if err != nil {
		return err
	}

	copy(b.data, buffer)

	b.cachedMeta = &meta
	return nil
}
