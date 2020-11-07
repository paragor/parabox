package engine

type Object struct {
	Key   string
	Value []byte
	Meta  map[string][]byte
}

type Query func(QueryResult) QueryResult

type QueryResult struct {
}

type StorageEngine interface {
	Insert(object Object) error
	Replace(object Object)
	FindByKey(key string) (*Object, error)
	Query(queries ...Query) ([]Object, error)
}
