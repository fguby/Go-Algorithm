package trees

type tree interface {
	IsEmpty() bool
	GetRootNode() interface{}
	Insert(e interface{})
	Containers(e interface{}) interface{}
}
