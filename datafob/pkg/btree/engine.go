package btree

import (
	"github.com/paragor/parabox/datafob/pkg/engine"
)

type BtreeEngine struct {
	data     []byte
	freeData int
	meta     []byte
	freeMeta int
}

func (e *BtreeEngine) Insert(object engine.Object) error {
	if e.freeData+len(object.Value) > len(e.data) {
		return engine.EndOfMemory
	}

	panic("implement me")
}

func (e *BtreeEngine) Replace(object engine.Object) {
	panic("implement me")
}

func (e *BtreeEngine) FindByKey(key string) (*engine.Object, error) {
	panic("implement me")
}

func (e *BtreeEngine) Query(queries ...engine.Query) ([]engine.Object, error) {
	panic("implement me")
}
