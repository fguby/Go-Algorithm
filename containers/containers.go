package containers

//Container是所有数据结构的基本接口
type Container interface {
	IsEmpty() bool
	GetSize() int
	Clear()
	Values() []interface{}
}

type IllegalError struct {
	s string
}

func (e *IllegalError) Error() string {
	return e.s
}
