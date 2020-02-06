package queue

type Queue interface {
	//是否为空
	IsEmpty() bool
	//获取size
	GetSize() int
	//入队
	Enqueue(e interface{}) error
	//出队
	Dequeue() (interface{}, error)
}
