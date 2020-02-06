package stack

type Stack interface {
	//
	GetSize() int
	//
	Resize() error
	//
	Push(e interface{}) error
	//
	Peek() (interface{}, error)
	//
	Pop() (interface{}, error)
}